<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import type { Rankings } from '@/types'
import VideoList from '@/components/VideoList.vue'

const rankings = ref<Rankings>([])
const isLoaded = ref(false)
const blockedOwnerIds = ref<Set<string>>(new Set())

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
    <ul class="rank">
      <li v-for="rank in filteredRankings[0]?.length" :key="rank">
        {{ rank }}
      </li>
    </ul>
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
  $rank-width: 50px;
  display: flex;

  .rank {
    width: $rank-width;
    list-style-type: none;
    margin: 0;
    padding: 0;

    li {
      height: 270px;
      padding: 3px;
      margin: 1px 0;
      border-bottom: 1px solid;
      font-size: 1.5rem;
      text-align: right;
    }
  }
}
</style>