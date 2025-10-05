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
  <ul class="video-list">
    <li v-for="video in ranking" :key="video.OwnerId" class="video-item">
      <div class="controls">
        <Slider :is_checked="video.IsMuted" @change="emit('toggleBlock', video.OwnerId)" />
      </div>
      <a
        v-show="!video.IsMuted && !video.IsPaymentRequired"
        :href="video.WatchURL"
        target="_blank"
        class="video-link"
      >
        <div class="thumbnail">
          <img :src="video.ThumbnailURL" loading="lazy" alt="thumbnail">
        </div>
        <p class="title">{{ video.Title }}</p>
      </a>
    </li>
  </ul>
</template>

<style lang="scss" scoped>
.video-list {
  width: calc((98vw - 3.47vw) / 4);
  list-style: none;
  margin: 0;
  padding: 0;
}

.video-item {
  padding: 0.56vw;
  margin: 0.56vw 0;
  border-bottom: 0.07vw solid var(--color-border);
  background: var(--color-background-soft);
  border-radius: 0.56vw;
  transition: transform 0.2s ease, box-shadow 0.2s ease;

  &:hover {
    transform: translateY(-0.14vw);
    box-shadow: 0 0.28vw 0.83vw rgba(0, 0, 0, 0.1);
  }
}

.controls {
  text-align: right;
  margin-bottom: 0.56vw;
}

.video-link {
  display: block;
  text-decoration: none;
  color: var(--color-text);

  &:visited {
    color: #8057b0;

    .thumbnail img {
      filter: grayscale(30%);
    }

    .title {
      opacity: 0.85;
      font-weight: 300;
      text-decoration: underline dotted;
    }
  }
}

.thumbnail {
  position: relative;
  width: 100%;
  padding-top: 56.25%;
  margin-bottom: 0.56vw;
  border-radius: 0.28vw;
  overflow: hidden;

  img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: opacity 0.2s ease;

    &:hover {
      opacity: 0.9;
    }
  }
}

.title {
  display: -webkit-box;
  -webkit-line-clamp: 4;
  line-clamp: 4;
  -webkit-box-orient: vertical;
  overflow: hidden;
  font-size: 0.9rem;
  line-height: 1.4;
  margin: 0;
  padding: 0 0.28vw;
}
</style>