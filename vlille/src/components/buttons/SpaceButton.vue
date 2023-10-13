<template>
  <div class="places-container available-place">
    <div class="icon">
      <IconParking v-if="free"></IconParking>
      <IconBike v-else></IconBike>
    </div>
    {{ badgeText }}
  </div>
</template>

<script lang="ts">
import { computed } from 'vue'
import IconParking from '../icons/IconParking.vue'
import IconBike from '../icons/IconBike.vue'
export default {
  props: {
    items: {
      type: Number,
      required: true
    },
    free: {
      type: Boolean,
      required: true
    }
  },
  setup(props) {
    const badgeText = computed(() => {
      if (props.free) {
        return `${props.items} places`
      } else {
        return `${props.items} velos`
      }
    })
    return {
      badgeText
    }
  },
  components: { IconParking, IconBike }
}
</script>
