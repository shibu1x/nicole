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
        <div class="thumbnail-wrapper">
          <img :src="video.ThumbnailURL" loading="lazy" alt="thumbnail">
        </div>
        <p>{{ video.Title }}</p>
      </a>
    </li>
  </ul>
</template>

<style lang="scss" scoped>
.main {
  width: calc((98vw - 50px) / 4);
  list-style-type: none;
  margin: 0;
  padding: 0;

  li {
    padding: 8px;
    margin: 8px 0;
    border-bottom: 1px solid var(--color-border);
    background: var(--color-background-soft);
    border-radius: 8px;
    transition: transform 0.2s ease, box-shadow 0.2s ease;

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }

    div {
      text-align: right;
      margin: 0 0 8px;
    }

    a {
      display: block;
      text-decoration: none;
      color: var(--color-text);

      &:visited {
        color: #8057b0; // より目立つ紫色に変更
        
        .thumbnail-wrapper {
          img {
            filter: grayscale(30%); // サムネイルをやや灰色がかった表示に
          }
        }

        p {
          opacity: 0.85;
          font-weight: 300; // フォントを細めに
          text-decoration: underline; // 下線を追加
          text-decoration-style: dotted; // 下線を点線に
        }
      }

      .thumbnail-wrapper {
        position: relative;
        width: 100%;
        padding-top: 56.25%; /* 16:9のアスペクト比 */
        margin-bottom: 8px;
        border-radius: 4px;
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

      p {
        display: -webkit-box;
        -webkit-line-clamp: 4;
        line-clamp: 4;
        -webkit-box-orient: vertical;
        overflow: hidden;
        font-size: 0.9rem;
        line-height: 1.4;
        margin: 0;
        padding: 0 4px;
      }
    }
  }
}
</style>