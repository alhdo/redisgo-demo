export interface ExtraInfo {
    address: string;
    city: string;
    last_update: string;
    online: boolean;
    'payment-terminal': boolean;
    status: string;
    uid: string;
  }
  
export  interface BikeStation {
    empty_slots: number;
    extra: ExtraInfo;
    free_bikes: number;
    id: string;
    latitude: number;
    longitude: number;
    name: string;
    timestamp: string;
  }

  export type BikeStationArray = BikeStation[];
