<script lang="ts">
	const { events } = $props();

	import { onMount } from 'svelte';
	import { ScheduleXCalendar } from '@schedule-x/svelte';
	import {
		createCalendar,
		createViewDay,
		createViewList,
		createViewMonthGrid,
		createViewWeek
	} from '@schedule-x/calendar';
	import '@schedule-x/theme-shadcn/dist/index.css';
	import 'temporal-polyfill/global';

	import { createResizePlugin } from '@schedule-x/resize';
	import { createDragAndDropPlugin } from '@schedule-x/drag-and-drop';
	import { createEventModalPlugin } from '@schedule-x/event-modal';

	let calendarApp: any = $state(null);

	onMount(() => {
		calendarApp = createCalendar({
			plugins: [createResizePlugin(), createDragAndDropPlugin(), createEventModalPlugin()],
			isDark: true,
			theme: 'shadcn',
			views: [createViewDay(), createViewWeek(), createViewMonthGrid(), createViewList()],
			events: events
		});
	});
</script>

{#if calendarApp}
	<ScheduleXCalendar {calendarApp} />
{/if}
