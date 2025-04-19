<script setup lang="ts">
defineProps<{
    checked: boolean
}>()
</script>

<template>
    <input type="checkbox" class="cursor-pointer" :checked />
</template>

<style scoped>
input[type="checkbox"] {
    accent-color: var(--p-primary-color);
}
</style>