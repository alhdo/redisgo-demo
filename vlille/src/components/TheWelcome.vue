<template>
  
  <div class="accordion" id="accordionExample">
  <div class="accordion-item">
    <h2 class="accordion-header" id="headingOne">
      <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
        Stations V'Lille
      </button>
    </h2>
    <div id="collapseOne" class="accordion-collapse collapse show" aria-labelledby="headingOne" data-bs-parent="#accordionExample">
      <div class="accordion-body">
        <div class="vlille-main-container">
    <VeloItem v-for="station in bikeStations" :key="station.id">
        <template #heading>{{ station.name }}</template>
        <template #address>{{ station.extra.address }}</template>
        <template #city>{{ station.extra.city }}</template>
        <template #empty-space>
          <SpaceButton :items=station.empty_slots :free=true></SpaceButton>
          <SpaceButton :items=station.free_bikes  :free=false></SpaceButton>
        </template>
    </VeloItem>

    
  </div>
      </div>
    </div>
  </div>


</div>
  
</template>

<script setup lang="ts">
import VeloItem from './VeloItem.vue'
import SpaceButton from './buttons/SpaceButton.vue'
import { useBikeStore} from '../stores/station';
import { storeToRefs } from 'pinia';

const { bikeStations } = storeToRefs(useBikeStore())
const { fetchBikeStations, initWebSocket } = useBikeStore()
fetchBikeStations()
initWebSocket()
</script>