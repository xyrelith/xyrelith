import type { PageLoad } from './$types';
import 'temporal-polyfill/global';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`http://xyrelith-api:2712/api/listEvents`);
	if (!res.ok) {
		throw new Error('Failed to fetch data');
	}

	interface EventResponse {
		events: Event[];
	}

	interface Event {
		id: number;
		title: string;
		priority: number;
		startDate: string;
		endDate: string;
	}

	interface CalendarEvent {
		id: string | number;
		start: Temporal.ZonedDateTime;
		end: Temporal.ZonedDateTime;
		title: string;
	}

	const eventsAPI: EventResponse = await res.json();

	function appendUTC(events: EventResponse): CalendarEvent[] {
		const eventsFormatted: CalendarEvent[] = events.events.map((event) => {
			return {
				id: `${event.id}`,
				title: `${event.title}`,
				end: Temporal.ZonedDateTime.from(`${event.endDate}[Europe/Warsaw]`),
				start: Temporal.ZonedDateTime.from(`${event.startDate}[Europe/Warsaw]`)
			};
		});

		return eventsFormatted;
	}

	console.log(appendUTC(eventsAPI));

	const events = appendUTC(eventsAPI);

	return { events };
};
