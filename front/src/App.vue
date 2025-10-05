<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useTitle } from '@vueuse/core'
import axios from 'axios'
import type { Rankings } from '@/types'
import VideoList from '@/components/VideoList.vue'

const rankings = ref<Rankings>([])
const isLoaded = ref(false)
const blockedOwnerIds = ref<Set<string>>(new Set())
const blockedOwnerIdsText = ref('')
const showBlockedControl = ref(false)
useTitle('NicoRanking')

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
  blockedOwnerIdsText.value = Array.from(blockedOwnerIds.value).join('\n')
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
  blockedOwnerIdsText.value = Array.from(blockedOwnerIds.value).join('\n')
}

const saveFromTextarea = () => {
  const ids = blockedOwnerIdsText.value
    .split('\n')
    .map(id => id.trim())
    .filter(id => id !== '')
  blockedOwnerIds.value = new Set(ids)
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
    <header>
      <button @click="showBlockedControl = !showBlockedControl">
        {{ showBlockedControl ? 'Hide' : 'Edit BlockList' }}
      </button>
      <div v-show="showBlockedControl" class="blocklist-editor">
        <textarea
          v-model="blockedOwnerIdsText"
          placeholder="Enter one Owner ID per line to block"
          rows="10"
        />
        <button @click="saveFromTextarea">Save</button>
      </div>
    </header>
    <main>
      <VideoList
        v-for="(ranking, index) in filteredRankings"
        :key="index"
        :ranking="ranking"
        @toggle-block="toggleBlockOwner"
      />
    </main>
  </div>
</template>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
}

header {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  margin: 0 2vw;

  button {
    margin-left: auto;
  }

  .blocklist-editor {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    margin-top: 0.56vw;

    textarea {
      width: 20vw;
      margin-bottom: 0.56vw;
    }
  }
}

main {
  display: flex;
  gap: 1.1vw;
  padding: 1.1vw;
  max-width: 100vw;
  margin: 0 auto;
}
</style>