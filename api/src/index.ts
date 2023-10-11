import cors from 'cors';
import express from 'express';
import http from 'http';
import cron from 'node-cron';
import {config} from '~/config'
import { StationsController } from './resources/stations/stations.controller';
import { UnknownRoutesHandler } from './middlewares/unknownRoutes.handler';
import { ExceptionHandler } from './middlewares/exceptions.handler';
import { createClient } from 'redis'
import { Server } from 'socket.io';

const arrayOfIdsToFilter = [
    "73b459427749807256c72d657d4143ba",
    "c58daefddd7ed6fd804ac34f0d649b09",
    "0e2c305fdfd9a22663f84d49db6e3a13",
    "c34c07922a0f6f7616c6916d9b1030b3", 
    "3cbc6a8f69efdd22de732363b94af38a"];

const app = express()

app.use(express.json())

app.use(cors())



const server = new http.Server(app)

const wss = new Server(server,{
    cors: {
      origin: "*"
    }
  });

  wss.on('connection', (socket) => {
    console.log(`${socket.id} connected.`);
  })
const client = createClient()
const subscriber = createClient()
client.on('error', err => console.log('Redis Client Error', err));
client.on("connect", () => {
    console.log(`Redis connection established`);
  });
client.connect().then(()=> {
    console.log("Connected")
})
subscriber.connect().then(()=> {
    console.log("Connected")
})

const clients = new Set();

wss.on('connection', (ws) => {
    clients.add(ws);

    ws.on('close', () => {
        clients.delete(ws);
    });
});

app.get('/fetch', (req, res) => {
    const job = {
        jid: 'test123',
        data: 'test'
    };

    client.publish('jobTopic', JSON.stringify(job))
    res.status(200).json({ message: 'Job sent successfully' })
});

subscriber.subscribe("jobResponseTopic", (message) => {
    console.log(message);
    const dataResponse = JSON.parse(message);

    client.get(dataResponse.taskID).then((response) => {
        const jsonData = JSON.parse(response as string)
        const filteredData = jsonData.filter((item: any)  => arrayOfIdsToFilter.includes(item.id))
    
        if (response !== null) {
            // wss.emit(response) 
            wss.send(JSON.stringify(filteredData))
            
        }
       
    });
})


app.use('/vlille', StationsController)

app.get('/', (req, res) => res.send('🏠'))

app.all('*', UnknownRoutesHandler)
app.use(ExceptionHandler)
// app.listen(config.API_PORT, () => console.log(`Listening on port ${config.API_PORT}`))
server.listen(config.API_PORT, () => console.log(`Listening on port ${config.API_PORT}`))

// cron.schedule('*/1 * * * *', async () => {
//     const job = {
//         jid: 'test123',
//         data: 'test'
//     };

//     client.publish('jobTopic', JSON.stringify(job))
//     console.log("Job submitted");
//   });