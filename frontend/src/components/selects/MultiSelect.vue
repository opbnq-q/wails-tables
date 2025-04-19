<script setup lang="ts" generic="T extends IEntity">
import { structView } from '../../utils/structs/structs-view.util';
import type { IEntity } from '../../types/entity.type';
import Checkbox from '../checkboxes/Checkbox.vue';

const { entityId } = defineProps<{
    options: T[]
    path?: string[]
    entityId: string
}>()

const selected = defineModel<T[]>({ default: [] })

const pushOrRemove = (option: T) => {
    if (selected.value.some(s => s.Id === option.Id)) {
        selected.value = selected.value.filter(s => s.Id !== option.Id)
    } else {
        selected.value.push(option)
    }
}

const setNullIds = () => {
    selected.value = selected.value.map(item => {
        (item as any)[entityId] = 0
        return item
    })
}

</script>

<template>
    <div class="">
        <ul class="max-h-48 h-auto overflow-y-auto background rounded-md p-3 w-full native-border">
            <li v-for="option in options" :key="option.Id" class="flex items-center gap-2">
                <Checkbox :checked="selected.some(item => item.Id == option.Id)" @click="pushOrRemove(option)" />
                <label :for="option.Id.toString()">{{ structView(option, path) }}</label>
            </li>
        </ul>
    </div>
</template>