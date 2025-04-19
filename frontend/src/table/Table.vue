<script setup lang="ts" generic="T extends IEntity">
import {onMounted, ref, watch, type UnwrapRef} from "vue";
import type {TableProps} from "../types/table-props.type";
import {DataTable, Column, Button, InputText} from "primevue";
import {manyStructsView} from "../utils/structs/structs-view.util";
import type {TableEmits} from "../types/table-emits.type";
import FloatingButton from "../components/buttons/FloatingButton.vue";
import type {IEntity} from "../types/entity.type";
import DialogWindow from "./DialogWindow.vue";
import {viewDate} from "../utils/date/view.util";
import SortingOptions from "./SortingOptions.vue";
import type {SortOptions} from "../types/sort-options.type.ts";

const props = defineProps<TableProps<T>>();

type Key = keyof T;
const keys = Object.keys(props.scheme) as Key[];
const emits = defineEmits<TableEmits>();

const showCreate = ref(false);
const createItem = ref<null | T>(null);

const showUpdate = ref(false);
const updateItem = ref<null | T>(null);

watch(showUpdate, (value) => {
  if (!value) {
    updateItem.value = null;
  }
});

watch(updateItem, (value) => {
  if (value) {
    showUpdate.value = true;
  } else {
    showUpdate.value = false;
  }
});

watch(showCreate, (value) => {
  if (value) {
    createItem.value = props.getDefaults();
  } else {
    createItem.value = null;
  }
});

watch(createItem, (value) => {
  if (value) {
    showCreate.value = true;
  } else {
    showCreate.value = false;
  }
});

const handleFloatingButtonClick = () => {
  emits('onCreateOpen')
  emits('onOpen')
  showCreate.value = true;
};
const slots = defineSlots();

const createSlotName = (key: any) => key + "Create";
const updateSlotName = (key: any) => key + "Update";

watch(() => props.items, () => {
  if (props.colorize) {
    const trs = document.querySelectorAll("tr");
    props.items.forEach(item => {
      const tr = trs[item.Id];
      if (tr) {
        tr.style.backgroundColor = props.colorize!(item);
      }
    })
  }
})

const input = ref('')
const sortOptions = defineModel<SortOptions<T>>('sort-options');
const showSortOptions = ref(false)

</script>

<template>
  <div class="m-2 flex flex-col items-center gap-2">
    <div class="flex flex-col gap-2 relative">
      <div class="flex items-center justify-center gap-2 h-10">
        <Button severity="secondary" class="h-full aspect-square" @click="showSortOptions = !showSortOptions">
          <span class="pi pi-sort"></span>
        </Button>

        <InputText class="h-full w-64" v-model="input"/>
        <Button severity="secondary" class="h-full aspect-square" @click="emits('onSearch', input)">
          <span class="pi pi-search"></span>
        </Button>
      </div>
      <SortingOptions v-model="sortOptions as SortOptions<T>" :scheme v-if="showSortOptions"
                      class="absolute z-10 w-full top-full translate-y-2" :load="props.load"></SortingOptions>
    </div>

  </div>
  <DialogWindow :name :load :items :validate :scheme :service :get-defaults v-model:item="<T>createItem"
                v-model:show="showCreate" @on-save="data => emits('onSave', data)"
                @on-save-create="data => emits('onSaveCreate', data)">
    <template v-for="key in keys" #[key]="{ data }">
      <slot :name="<string>key" :data="data"></slot>
    </template>
    <template v-for="key in keys" #[createSlotName(key)]="{ data }">
      <slot :name="createSlotName(key)" :data="data"></slot>
    </template>
  </DialogWindow>
  <DialogWindow :name :load :items :validate :scheme update-mode :service :get-defaults v-model:item="<T>updateItem"
                v-model:show="showUpdate" @on-save="data => emits('onSave', data)"
                @on-save-update="data => emits('onSaveUpdate', data)">
    <template v-for="key in keys" #[key]="{ data }">
      <slot :name="<string>key" :data="data"></slot>
    </template>
    <template v-for="key in keys" #[updateSlotName(key)]="{ data }">
      <slot :name="updateSlotName(key)" :data="data"></slot>
    </template>
  </DialogWindow>
  <div>
    <DataTable :value="<[]>items" paginator :rows="10">
      <template #header v-if="props.name">
        <p>{{ props.name }}</p>
      </template>
      <template v-for="key in keys">
        <Column :header="props.scheme[key]?.russian" v-if="!props.scheme[key].hidden">
          <template #body="{ data }">
            <p class="truncate max-w-56" v-tooltip="viewDate(manyStructsView(
                            data[key],
                            props.scheme[key]?.type?.nested?.field,
                        ), scheme[key]?.type?.primitive)">
              {{
                viewDate(manyStructsView(
                    data[key],
                    props.scheme[key]?.type?.nested?.field,
                ), scheme[key]?.type?.primitive)
              }}
            </p>
          </template>
        </Column>
      </template>
      <Column header="Действия">
        <template #body="{ data }">
          <div class="flex gap-2">
            <Button severity="secondary" icon="pi pi-pencil" @click="async () => {
                            await emits('onUpdateOpen')
                            await emits('onOpen')
                            updateItem = data
                        }"></Button>
            <Button severity="danger" icon="pi pi-trash" @click="async () => {
                            await emits('onDelete', data)
                            await props.service.delete(data.Id)
                            await load()
                        }"></Button>
          </div>
        </template>
      </Column>
      <template #paginatorstart>
        <Button severity="secondary" @click="props.load">
          <span class="pi pi-refresh"></span>
        </Button>
      </template>
      <template #paginatorend>
        <span></span>
      </template>
    </DataTable>
    <FloatingButton @click="handleFloatingButtonClick"/>
  </div>
</template>
