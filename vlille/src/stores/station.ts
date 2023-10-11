import { defineStore } from "pinia";
import {type BikeStationArray } from '../models/station.model'
import { socket } from "@/socket";
export const useBikeStore = defineStore({
    id: 'bike',
    state: () => ({
        bikeStations: [] as BikeStationArray,
    }),
    getters: {
        getStations(state) {
            return state.bikeStations;
        }
    },
    actions: {
        async fetchBikeStations() {
            const response = await fetch('http://127.0.0.1:3000/fetch');
            const data = await response.json();
            console.log(data);
            // this.bikeStations = data;
        },
        initWebSocket() {
            console.log("Web socket inited");
            // const ws = new WebSocket('ws://localhost:3000');
           
            // ws.on('message', (message: string) => {
            //   // Handle WebSocket message here and update the store
            //   const data = JSON.parse(message);
            //   this.bikeStations = data;
            // });
            socket.on('message', (message: any) => {
                
                const data = JSON.parse(message)
                console.log(data);
                // this.bikeStations = [...data];
                this.$state.bikeStations = [...data];
            });
          }
    }
})