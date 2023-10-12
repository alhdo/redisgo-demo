import { defineStore } from "pinia";
import {type BikeStationArray } from '../models/station.model'
import { socket } from "@/socket";
export const useBikeStore = defineStore({
    id: 'bike',
    state: () => ({
        bikeStations: [] as BikeStationArray,
        loading: false,
        searchTerm: '',
    }),
    getters: {
        getStations(): any {
            const query = this.searchTerm.toLowerCase()
            return this.bikeStations.filter((item) => item.name.toLowerCase().includes(query));
        }
    },
    actions: {
        async fetchBikeStations() {
            this.loading = true;

            const response = await fetch('http://127.0.0.1:3000/fetch');
            this.loading = false;
            await response.json();
        },
        // initWebSocket() {
        //     console.log("Web socket inited");
        //     // const ws = new WebSocket('ws://localhost:3000');
           
        //     // ws.on('message', (message: string) => {
        //     //   // Handle WebSocket message here and update the store
        //     //   const data = JSON.parse(message);
        //     //   this.bikeStations = data;
        //     // });
        //     socket.on('message', (message: any) => {
                
        //         const data = JSON.parse(message)
        //         console.log(data);
        //         // this.bikeStations = [...data];
        //         this.$state.bikeStations = [...data];
        //     });
        //   },

          setSearchTerm(query: string) {
            this.searchTerm = query;
          }
    }
})