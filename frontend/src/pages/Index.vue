
<!--<script setup lang="ts">-->
<!--import { Button } from 'primevue';-->
<!--import NavCard from '../components/cards/NavCard.vue';-->
<!--import { ExportAllEntities, ImportAllEntities } from '../../bindings/app/internal/services/excelmodule';-->
<!--import { useErrorStore } from '../stores/error.store';-->

<!--const errorStore = useErrorStore()-->

<!--const importFromExcel = async () => {-->
<!--  try {-->
<!--    await ImportAllEntities()-->
<!--  } catch(e) {-->
<!--    errorStore.summon((e as Error).message)-->
<!--  }-->
<!--}-->

<!--const exportFromExcel = async () => {-->
<!--  try {-->
<!--    await ExportAllEntities()-->
<!--  } catch(e) {-->
<!--    errorStore.summon((e as Error).message)-->
<!--  }-->
<!--}-->
<!--</script>-->

<!--<template>-->
<!--  <div class="flex w-screen h-screen items-center gap-5 justify-center">-->
<!--    <NavCard :title="'Пользователь'" :to="'/user'" :content="'Не может пользоваться туалетной бумагой'" />-->
<!--  </div>-->
<!--  <footer class="fixed w-full bottom-10 flex items-center justify-center gap-2">-->
<!--    <Button severity="secondary" @click="importFromExcel">Импортировать данные</Button>-->
<!--    <Button severity="secondary" @click="exportFromExcel">Экспортировать данные</Button>-->
<!--  </footer>-->
<!--</template>-->

<template>
</template>
<script setup lang="ts">
</script>