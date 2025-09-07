<script lang="ts">
	import Calendar from '../components/calendar.svelte';

	import type { PageProps } from './$types';
	let { data }: PageProps = $props();

	console.log(data.events);
</script>

<div class="calendar-header"></div>
<Calendar events={data.events} />
