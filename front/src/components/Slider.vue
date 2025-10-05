<script setup lang="ts">
defineProps<{
  is_checked: boolean
}>()

const emit = defineEmits(['change'])
</script>

<template>
  <label class="switch">
    <input type="checkbox" :checked="is_checked" @change="emit('change')">
    <span class="slider round" />
  </label>
</template>

<style lang="scss" scoped>
$switch-width: 2.78vw;
$switch-height: 1.39vw;
$slider-size: 1.11vw;
$slider-offset: 0.14vw;

.switch {
  position: relative;
  display: inline-block;
  width: $switch-width;
  height: $switch-height;
  margin: 0.28vw;

  input {
    opacity: 0;
    width: 0;
    height: 0;

    &:checked + .slider {
      background-color: #4caf50;

      &::before {
        transform: translateX($switch-width - $slider-size - 0.28vw);
      }
    }

    &:focus + .slider {
      box-shadow: 0 0 0.07vw #4caf50;
    }
  }
}

.slider {
  position: absolute;
  cursor: pointer;
  inset: 0;
  background-color: #ccc;
  transition: 0.3s ease-in-out;

  &::before {
    position: absolute;
    content: '';
    height: $slider-size;
    width: $slider-size;
    left: $slider-offset;
    bottom: $slider-offset;
    background-color: white;
    transition: 0.3s ease-in-out;
    box-shadow: 0 0.14vw 0.28vw rgba(0, 0, 0, 0.2);
  }

  &.round {
    border-radius: $switch-height;

    &::before {
      border-radius: 50%;
    }
  }
}
</style>