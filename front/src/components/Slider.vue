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
/* The switch - the box around the slider */
$switch_width: 60px;
$slider_width: 26px;

label.switch {
    position: relative;
    display: inline-block;
    width: $switch_width;
    height: $switch_width - $slider_width;

    /* Hide default HTML checkbox */
    input {
        opacity: 0;
        width: 0;
        height: 0;

        &:checked+.slider {
            background-color: #2196F3;

            &:before {
                transform: translateX($slider_width);
            }
        }

        &:focus+.slider {
            box-shadow: 0 0 1px #2196F3;
        }
    }
}

/* The slider */
span.slider {
    position: absolute;
    cursor: pointer;
    inset: 0;
    background-color: #ccc;
    transition: .4s;

    &:before {
        position: absolute;
        content: "";
        height: $slider_width;
        width: $slider_width;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: .4s;
    }

    /* Rounded sliders */
    &.round {
        border-radius: $switch_width - $slider_width;

        &:before {
            border-radius: 50%;
        }
    }
}
</style>