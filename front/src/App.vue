<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useTitle } from '@vueuse/core'
import axios from 'axios'
import type { Rankings } from '@/types'
import VideoList from '@/components/VideoList.vue'

const rankings = ref<Rankings>([])
const isLoaded = ref(false)
const blockedOwnerIds = ref<Set<string>>(new Set())
useTitle('ニコニコ動画ランキング')

const filteredRankings = computed(() => {
  if (!isLoaded.value) return []

  return rankings.value.map(ranking => 
    ranking.map(video => ({
      ...video,
      IsMuted: blockedOwnerIds.value.has(video.OwnerId)
    }))
  )
})

const loadBlockedOwnerIds = () => {
  const saved = localStorage.getItem('block_owner_ids')
  blockedOwnerIds.value = new Set(saved ? JSON.parse(saved) : [])
}

const saveBlockedOwnerIds = () => {
  localStorage.setItem(
    'block_owner_ids',
    JSON.stringify(Array.from(blockedOwnerIds.value))
  )
}

const toggleBlockOwner = (ownerId: string) => {
  const ids = blockedOwnerIds.value
  ids.has(ownerId) ? ids.delete(ownerId) : ids.add(ownerId)
  saveBlockedOwnerIds()
}

onMounted(async () => {
  try {
    const response = await axios.get('https://shibu1x-public.s3.ap-northeast-1.amazonaws.com/test.json')
    rankings.value = response.data
    loadBlockedOwnerIds()
    isLoaded.value = true
  } catch (error) {
    console.error('Failed to load rankings:', error)
  }
})
</script>

<template>
  <div class="container">
    <VideoList
      v-for="(ranking, index) in filteredRankings"
      :key="index"
      :ranking="ranking"
      @toggle-block="toggleBlockOwner"
    />
  </div>
</template>

<style lang="scss" scoped>
.container {
  display: flex;
  gap: 16px;
  padding: 16px;
  max-width: 1440px;
  margin: 0 auto;
}

.title-container {
  width: 100%;
  text-align: center;
  margin-bottom: 16px;
}

.title-input {
  font-size: 1.5rem;
  padding: 8px 16px;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-background-soft);
  color: var(--color-text);
  width: 100%;
  max-width: 400px;
  text-align: center;
  
  &:focus {
    outline: none;
    border-color: var(--color-border-hover);
  }
}
</style>