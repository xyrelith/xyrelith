<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import DarkModeButton from '$lib/components/DarkModeButton.svelte';
</script>

<Sidebar.Root>
	<Sidebar.Header />
	<Sidebar.Content>
		<DarkModeButton />
		<Sidebar.Group />
		<Sidebar.Group />
	</Sidebar.Content>
	<Sidebar.Footer />
</Sidebar.Root>
