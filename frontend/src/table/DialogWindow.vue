<script setup lang="ts" generic="T extends IEntity">
import { Button, DatePicker, Dialog, InputNumber, InputText, Select, Textarea, ToggleSwitch } from 'primevue';
import type { IEntity } from '../types/entity.type';
import type { Scheme } from '../types/scheme.type';
import type { IService } from '../types/service.type';
import { manyStructsView } from '../utils/structs/structs-view.util';
import { toDate, toTimestamp } from '../utils/date/converters.util';
import MultiSelect from '../components/selects/MultiSelect.vue';
import type { Validate } from '../types/validate.type';
import { useErrorStore } from '../stores/error.store';

const showCreate = defineModel<boolean>('show')
const editableItem = defineModel<T>('item')

const errorStore = useErrorStore()

const props = defineProps<{
    scheme: Scheme<T>,
    getDefaults: () => Partial<T>,
    service: IService<T>,
    updateMode?: boolean
    validate: Validate<T>
    items: T[]
    load(): Promise<T[]>
    name?: string
}>()

type Key = keyof T
const keys = Object.keys(props.scheme) as Key[]

const emits = defineEmits<{
    (e: 'onSave', data: T): Promise<void> | void
    (e: 'onSaveUpdate', data: T): Promise<void> | void
    (e: 'onSaveCreate', data: T): Promise<void> | void
}>()

async function handleSave() {
    const mode = props.updateMode ? 'update' : 'create';
    const result = await props.validate(editableItem.value as T, mode);
    if (result.status === 'error') {
        errorStore.summon(result.message ?? 'Ошибка валидации');
        return;
    }
    try {
        if (props.updateMode) {
            await props.service.update.call(props.service, editableItem.value as T);
            await emits('onSaveUpdate', editableItem.value as T);
            await emits('onSave', editableItem.value as T);
        } else {
            await props.service.create.call(props.service, editableItem.value as T);
            await emits('onSaveCreate', editableItem.value as T);
            await emits('onSave', editableItem.value as T);
        }
    } catch (e) {
        errorStore.summon((e as Error).message ?? 'Ошибка сохранения');
        return
    }
    props.load()
    showCreate.value = false;
}

</script>

<template>
    <Dialog v-model:visible="showCreate" class="w-[500px]">
        <template #header>
            <h1>{{ props.updateMode ? 'Изменить' : 'Создать' }} {{ props.name?.toLowerCase() }}</h1>
        </template>
        <div class="flex flex-col justify-start gap-5 min-h-[40vh]">
            <div v-for="key in keys" v-show="!props.scheme[key].hidden && !props.scheme[key].readonly"
                class="flex items-center gap-5">
                <h1 class="w-[200px]">{{ props.scheme[key].russian }}</h1>
                <div v-if="props.scheme[key]?.customWindow?.[props.updateMode ? 'update' : 'create']">
                    <slot :name="<string>key + (props.updateMode ? 'Update' : 'Create')" :data="editableItem"></slot>
                </div>
                <div v-else-if="props.scheme[key]?.customWindow?.common">
                    <slot :name="<string>key" :data="editableItem"></slot>
                </div>
                <InputNumber class="w-[300px]" v-model:model-value="<number>editableItem![key]"
                    v-else-if="props.scheme[key]?.type?.primitive === 'number'" />
                <InputText class="w-[300px]" v-model:model-value="<string>editableItem![key]"
                    v-else-if="props.scheme[key].type?.primitive === 'string'" />
                <DatePicker class="w-[300px]" :default-value="toDate(editableItem![key] as number)" @value-change="v => {
                    editableItem![key] = toTimestamp(v as Date) as any
                }" show-time v-else-if="props.scheme[key].type?.primitive === 'date'" />
                <Textarea class="w-[300px]" v-model="<string>editableItem![key]"
                    v-else-if="props.scheme[key].type?.primitive === 'multiple'" />
                <ToggleSwitch class="w-[300px]" v-model:model-value="<boolean>editableItem![key]"
                    v-else-if="props.scheme[key].type?.primitive === 'boolean'" />
                <Select v-else-if="props.scheme[key].type?.nested?.values && !props.scheme[key]?.many"
                    v-model:model-value="editableItem![key]" :options="props.scheme[key].type.nested.values"
                    :placeholder="`Выберите ${props.scheme[key].russian}`" class="w-[300px]">
                    <template #option="{ option }">
                        {{ manyStructsView(option, props.scheme[key].type.nested.field) }}
                    </template>
                    <template #value="{ value }">
                        {{ manyStructsView(value, props.scheme[key].type.nested.field) }}
                    </template>
                </Select>
                <MultiSelect v-else-if="props.scheme[key].type?.nested?.values && props.scheme[key]?.many"
                    class="w-[300px]" v-model="<T[]>editableItem![key]" :options="props.scheme[key].type.nested.values"
                    :path="props.scheme[key].type.nested.field" :entity-id="props.scheme.entityId" />
            </div>
        </div>
        <template #footer>
            <Button class="mt-5" @click="handleSave">Сохранить</Button>
        </template>
    </Dialog>
</template>