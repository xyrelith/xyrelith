<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';

	import { ModeWatcher, setMode } from 'mode-watcher';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/AppSidebar.svelte';
	import { onMount } from 'svelte';

	let { children } = $props();

	onMount(() => {
		setMode('dark');
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<ModeWatcher />
<Sidebar.Provider>
	<AppSidebar />
	<main>
		{@render children?.()}
	</main>
</Sidebar.Provider>

<style>
	main {
		width: 100%;
	}
</style>
