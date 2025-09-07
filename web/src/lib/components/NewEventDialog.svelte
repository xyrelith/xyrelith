<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import PrioritySelect from './PrioritySelect.svelte';
	import TimeDatePicker from './TimeDatePicker.svelte';
</script>

<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'outline' })}>New Event</Dialog.Trigger>
	<Dialog.Content class="items-center sm:max-w-[425px]">
		<Dialog.Header>
			<Dialog.Title>New event</Dialog.Title>
			<Dialog.Description>Fill info about your new event.</Dialog.Description>
		</Dialog.Header>
		<div class="grid gap-4 py-4">
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="title" class="text-right">Title</Label>
				<Input id="title" placeholder="My new event" class="col-span-3" />
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="start-date" class="text-right">Start Date</Label>
				<TimeDatePicker />
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="end-date" class="text-right">End Date</Label>
				<TimeDatePicker />
			</div>
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="priority" class="text-right">Priority</Label>
				<PrioritySelect />
			</div>
		</div>
		<Dialog.Footer>
			<Button type="submit">Create</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
