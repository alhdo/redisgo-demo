import { mount } from '@vue/test-utils';
import TheWelcome from '../TheWelcome.vue';
import VeloItem from '../VeloItem.vue';
import { describe, expect, it, vi } from 'vitest';
import { createTestingPinia } from '@pinia/testing'
import { useBikeStore } from '@/stores/station';
import { createPinia } from 'pinia';

describe('TheWelcome', () => {
    it('renders loading text when loading is true', async () => {
        const wrapper = mount(TheWelcome, {
            global: {
                plugins: [createPinia()]
            }
        });
        const store = useBikeStore()
        store.loading = true;
        store.bikeStations = [];
        // Assert that the "Loading..." text is displayed
        expect(wrapper.text()).toContain('Loading...');
      });

      it('Store search term get called', async () => {
        const stations = [
            {
                "empty_slots": 18,
                "extra": {
                    "address": "Parvis Rotterdam",
                    "city": "LILLE",
                    "last_update": "2023-10-12T21:16:10+00:00",
                    "online": true,
                    "payment-terminal": true,
                    "status": "EN SERVICE",
                    "uid": "160"
                },
                "free_bikes": 2,
                "id": "3d91e5c7edcdf745a745ae95b3cf3721",
                "latitude": 50.637709,
                "longitude": 3.077478,
                "name": "PARVIS ROTTERDAM",
                "timestamp": "2023-10-12T21:26:33.309000Z"
            },
            {
                "empty_slots": 13,
                "extra": {
                    "address": "ALLEE VICTOR BASCH",
                    "city": "LA MADELEINE",
                    "last_update": "2023-10-12T21:16:10+00:00",
                    "online": true,
                    "payment-terminal": true,
                    "status": "EN SERVICE",
                    "uid": "96"
                },
                "free_bikes": 14,
                "id": "93051bcb05162e8baae7124c1385f2b8",
                "latitude": 50.64543,
                "longitude": 3.075116,
                "name": "ROMARIN",
                "timestamp": "2023-10-12T21:26:33.308000Z"
            }];
    
        const wrapper = mount(TheWelcome, {
            global: {
                plugins: [createTestingPinia(
                    {
                        createSpy: vi.fn,
                    }
                )]
            }
        });
        const store = useBikeStore()
        store.setSearchTerm('ROMARIN')
        expect(store.setSearchTerm).toHaveBeenCalledTimes(1)
      });

      it('renders stations when loading is false and there are stations', async () => {
        const stations = [
            {
                "empty_slots": 18,
                "extra": {
                    "address": "Parvis Rotterdam",
                    "city": "LILLE",
                    "last_update": "2023-10-12T21:16:10+00:00",
                    "online": true,
                    "payment-terminal": true,
                    "status": "EN SERVICE",
                    "uid": "160"
                },
                "free_bikes": 2,
                "id": "3d91e5c7edcdf745a745ae95b3cf3721",
                "latitude": 50.637709,
                "longitude": 3.077478,
                "name": "PARVIS ROTTERDAM",
                "timestamp": "2023-10-12T21:26:33.309000Z"
            },
            {
                "empty_slots": 13,
                "extra": {
                    "address": "ALLEE VICTOR BASCH",
                    "city": "LA MADELEINE",
                    "last_update": "2023-10-12T21:16:10+00:00",
                    "online": true,
                    "payment-terminal": true,
                    "status": "EN SERVICE",
                    "uid": "96"
                },
                "free_bikes": 14,
                "id": "93051bcb05162e8baae7124c1385f2b8",
                "latitude": 50.64543,
                "longitude": 3.075116,
                "name": "ROMARIN",
                "timestamp": "2023-10-12T21:26:33.308000Z"
            }];
    
        const wrapper = mount(TheWelcome, {
            global: {
                plugins: [createTestingPinia({
                    initialState: {
                        bike: {
                            bikeStations: stations,
                            loading: false
                        },
                        
                    },
                    createSpy: vi.fn
                })]
            }
        });
        console.log(wrapper.text())

        for (const station of stations) {
            expect(wrapper.text()).toContain(station.name);
          }
      });
})