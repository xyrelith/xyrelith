<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js';

	const priorities = [
		{ value: 'Low', label: 'Low' },
		{ value: 'Medium', label: 'Medium' },
		{ value: 'High', label: 'High' },
		{ value: 'Critical', label: 'Critical' }
	];

	let value = $state('');

	const triggerContent = $derived(
		priorities.find((f) => f.value === value)?.label ?? 'Select priority'
	);
</script>

<Select.Root type="single" name="priority" bind:value>
	<Select.Trigger id="priority" class="col-span-3 w-auto">
		{triggerContent}
	</Select.Trigger>
	<Select.Content>
		<Select.Group>
			<Select.Label>Priority</Select.Label>
			{#each priorities as priority (priority.value)}
				<Select.Item value={priority.value} label={priority.label}>
					{priority.label}
				</Select.Item>
			{/each}
		</Select.Group>
	</Select.Content>
</Select.Root>
