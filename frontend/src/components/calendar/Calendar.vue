<script setup lang="ts">
import { Qalendar } from "qalendar";
import { CalendarEvent } from "./event.class";

const config = {
  month: {
    showTrailingAndLeadingDates: false,
  },
  locale: 'ru-RU',
  defaultMode: 'month',
  isSilent: true,
};

defineProps<{
  events: CalendarEvent[]
}>()
</script>

<template>
  <Qalendar :config :events></Qalendar>
</template>

<style>
 @import "qalendar/dist/style.css";
</style>