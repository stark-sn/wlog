const weekElement = document.getElementById("week");
const errorElement = document.getElementById("error");
const dataElement = document.getElementById("data");

const dataDir = "data";

// State

let year;
let week;

// Helpers

/**
 * Create title.
 *
 * @param {string} year - Year
 * @param {string} week - Week
 * @returns {string} Title
 */
function title(year, week) {
	return `${year}-W${week}`;
}

/**
 * Format moment date object accorindg to application datetime format.
 *
 * @param {*} d - moment date object
 * @returns {string} Formatted datetime
 */
function format(d) {
	return d.format("HH:mm:ss");
}

// Data

/**
 * Load data and trigger rendering.
 */
function load() {

	const file = `${dataDir}/${year}/${week}.json`;
	weekElement.innerHTML = title(year, week);

	fetch(file).then(r => {
		if (r.ok) {
			return r.json();
		} else {
			throw new Error(`${r.status} - ${r.statusText} (${file})`);
		}
	}).then(renderData).catch(renderError);
}

// Controls

window.onpopstate = onPopState;

/**
 * React on history state change.
 *
 * @param {PopStateEvent} event - History event
 */
function onPopState(event) {
	year = event.state.year;
	week = event.state.week;
	load();
}

/**
 * Adjust selected week with provided function.
 *
 * @param {function(*): *} func - Control function that adjusts the provided moment object
 */
function changeWeek(func) {
	const d = func(moment(`${year} ${week}`, "GGGG WW"));
	year = d.format("GGGG");
	week = d.format("WW");

	history.pushState({
		year,
		week
	},
	`SST WLOG - ${title(year, week)}`,
	`?year=${year}&week=${week}`);

	load();
}

/**
 * Select current/last week.
 */
function currWeek() {
	changeWeek(() => moment().subtract(1, "week"));
}

/**
 * Select previous week.
 */
function prevWeek() {
	changeWeek(d => d.subtract(1, "week"));
}

/**
 * Select next week.
 */
function nextWeek() {
	changeWeek(d => d.add(1, "week"));
}

// Rendering

/**
 * Reset view.
 */
function reset() {
	dataElement.innerHTML = "";
	errorElement.innerHTML = "";
}

/**
 * Render wlog information.
 *
 * @param {*} data - wlog data
 */
function renderData(data) {

	reset();

	for (const date of Object.keys(data.Days)) {
		const day = data.Days[date];

		dataElement.innerHTML += `${date}\n`;

		if (day.Activities && day.Activities.length > 0) {
			let d = 0;
			dataElement.innerHTML += "\n";
			for (const act of day.Activities) {
				const start = moment(act.Start);
				const end = moment(act.End);
				const dur = end.diff(start);
				d += dur;
				dataElement.innerHTML += `\t${act.Title}:\t\t${format(start)} - ${format(end)}\t${format(moment.utc(dur))}\n`;
			}

			const dur = moment.utc(moment.duration(d).asMilliseconds());
			dataElement.innerHTML += `\n\t\t\t\t\t\t${format(dur)}\n\n`;
		}

		let breakTime = 0;
		if (day.Breaks && day.Breaks.length > 0 ) {
			dataElement.innerHTML += "\tBreaks:\n";
			for (const b of day.Breaks) {
				const start = moment(b.Start);
				const end = moment(b.End);
				const dur = end.diff(start);

				breakTime += dur;
				dataElement.innerHTML += `\t\t${format(start)} - ${format(end)}\t${format(moment.utc(dur))}\n`;
			}
		}


		const breakDur = moment.utc(moment.duration(breakTime).asMilliseconds());

		if (breakTime > 0) {
			dataElement.innerHTML += `\n\t\t\t\t\t${format(breakDur)}\n\n`;
		}

		let d = 0;
		for (const span of day.Spans) {
			const start = moment(span.Start);
			const end = moment(span.End);
			const dur = end.diff(start);
			d += dur;
			dataElement.innerHTML += `\t${format(start)} - ${format(end)}\t${format(moment.utc(dur))}\n`;
		}
		const workTime = d - breakTime;
		const dur = moment.utc(moment.duration(d).asMilliseconds());
		const workDur = moment.utc(moment.duration(workTime).asMilliseconds());
		dataElement.innerHTML += `\n\t\t\t\t${format(dur)}\n`;
		dataElement.innerHTML += `\t\t\t&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;-${format(breakDur)} (breaks)\n`;
		dataElement.innerHTML += `\t\t\t\t${format(workDur)}\n\n`;
	}
}

/**
 * Render error information.
 *
 * @param {*} e - Error
 */
function renderError(e) {
	reset();
	errorElement.innerHTML = e;
}

// Main

(function() {

	// Check if year/week are present in URL and load data.
	// If missing, load current/last week.
	const url = new URL(location.href);
	year = url.searchParams.get("year");
	week = url.searchParams.get("week");

	if (!(year && week)) {
		currWeek();
		return
	}

	load();
})();

