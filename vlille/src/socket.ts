import { reactive } from "vue";
import { io } from "socket.io-client";
import { useBikeStore} from './stores/station';

export const state = reactive({
    connected: false,
});
const URL = process.env.NODE_ENV === "production" ? undefined : "http://localhost:3000";
export const socket = io("http://api:3000");

socket.on("connect", () => {
  state.connected = true;
  console.log("Socket connected");
});

socket.on("disconnect", () => {
  state.connected = false;
});

socket.on("message", (message) => {
    const bikeStore = useBikeStore();
    bikeStore.bikeStations = JSON.parse(message);
})