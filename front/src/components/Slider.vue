<script setup lang="ts">

const props = defineProps<{
  is_checked: boolean
}>()

const emits = defineEmits(['change'])

const change = (value?: String): void => {
    emits('change', value)
}

</script>

<template>
    <label class="switch">
        <input type="checkbox" :checked="is_checked" @change="change()">
        <span class="slider round"></span>
    </label>
</template>

<style lang="scss" scoped>
$switch_width: 40px;
$switch_height: 20px;
$slider_size: 16px;

label.switch {
    position: relative;
    display: inline-block;
    width: $switch_width;
    height: $switch_height;
    margin: 4px;

    input {
        opacity: 0;
        width: 0;
        height: 0;

        &:checked + .slider {
            background-color: #4CAF50;

            &:before {
                transform: translateX($switch_width - $slider_size - 4px);
            }
        }

        &:focus + .slider {
            box-shadow: 0 0 1px #4CAF50;
        }
    }
}

.slider {
    position: absolute;
    cursor: pointer;
    inset: 0;
    background-color: #ccc;
    transition: .3s ease-in-out;

    &:before {
        position: absolute;
        content: "";
        height: $slider_size;
        width: $slider_size;
        left: 2px;
        bottom: 2px;
        background-color: white;
        transition: .3s ease-in-out;
        box-shadow: 0 2px 4px rgba(0,0,0,0.2);
    }

    &.round {
        border-radius: $switch_height;

        &:before {
            border-radius: 50%;
        }
    }
}
</style>