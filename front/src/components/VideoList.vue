<script setup lang="ts">
import type { Video } from '@/types'
import Slider from './Slider.vue'

defineProps<{
  ranking: Video[]
}>()

const emit = defineEmits<{
  'toggleBlock': [ownerId: string]
}>()
</script>

<template>
  <ul class="main">
    <li v-for="video in ranking" :key="video.OwnerId">
      <div>
        <Slider 
          :is_checked="video.IsMuted" 
          @change="emit('toggleBlock', video.OwnerId)" 
        />
      </div>
      <a 
        :href="video.WatchURL" 
        target="_blank" 
        v-show="!video.IsMuted && !video.IsPaymentRequired"
      >
        <img :src="video.ThumbnailURL" loading="lazy" alt="thumbnail">
        <p>{{ video.Title }}</p>
      </a>
    </li>
  </ul>
</template>

<style lang="scss" scoped>
.main {
  width: calc((98vw - 50px) / 5);
  list-style-type: none;
  margin: 0;
  padding: 0;

  li {
    height: 270px;
    padding: 3px;
    margin: 1px 0;
    border-bottom: 1px solid;

    div {
      text-align: right;
      margin: 0 0 5px;
    }

    a {
      img {
        display: block;
        max-width: 100%;
        max-height: 103px;
        height: auto;
        margin: 0 auto;
      }

      p {
        display: -webkit-box;
        -webkit-line-clamp: 4;
        line-clamp: 4;
        -webkit-box-orient: vertical;
        overflow: hidden;
      }
    }
  }
}
</style>