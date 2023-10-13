<template>
  <div class="accordion" id="accordionExample">
    <div class="accordion-item">
      <h2 class="accordion-header" id="headingOne">
        <button
          class="accordion-button"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#collapseOne"
          aria-expanded="true"
          aria-controls="collapseOne"
        >
          Stations V'Lille
        </button>
      </h2>
      <div
        id="collapseOne"
        class="accordion-collapse collapse show"
        aria-labelledby="headingOne"
        data-bs-parent="#accordionExample"
      >
        <div class="accordion-body">
          <div class="input-group mb-3">
            <input
              v-model="searchTerm"
              type="text"
              class="form-control"
              placeholder="Search..."
              aria-label="Search"
              id="search"
            />
          </div>

          <div class="vlille-main-container">
            <div class="no-result" v-if="!getStations.length && !loading">
              <EmptyView></EmptyView>
            </div>

            <VeloItem
              v-else-if="getStations.length && !loading"
              v-for="station in getStations"
              :key="station.id"
            >
              <template #heading>{{ station.name }}</template>
              <template #address>{{ station.extra.address }}</template>
              <template #city>{{ station.extra.city }}</template>
              <template #empty-space>
                <SpaceButton :items="station.empty_slots" :free="true"></SpaceButton>
                <SpaceButton :items="station.free_bikes" :free="false"></SpaceButton>
              </template>
            </VeloItem>

            <div class="loading-container" v-if="loading">
              <strong>Loading...</strong>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import VeloItem from './VeloItem.vue'
import SpaceButton from './buttons/SpaceButton.vue'
import EmptyView from './EmptyView.vue'
import { useBikeStore } from '../stores/station'
import { storeToRefs } from 'pinia'

const { loading, searchTerm, getStations } = storeToRefs(useBikeStore())
const { fetchBikeStations } = useBikeStore()
fetchBikeStations()
// initWebSocket()
</script>
