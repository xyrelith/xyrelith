<script lang="ts">
	import NewEventDialog from '@/components/NewEventDialog.svelte';
	import Calendar from '../components/calendar.svelte';

	import type { PageProps } from './$types';
	import SidebarTrigger from '@/components/ui/sidebar/sidebar-trigger.svelte';

	let { data }: PageProps = $props();
</script>

<div class="m-4 flex flex-row items-center">
	<SidebarTrigger class="mr-4 scale-150" />
	<NewEventDialog />
</div>
<Calendar events={data.events} />
