const weekElement = document.getElementById("week");
const errorElement = document.getElementById("error");
const dataElement = document.getElementById("data");

const dataDir = "data";

let year;
let week;

async function loadFile() {

	const file = `${dataDir}/${year}/${week}.json`;

	return await fetch(file).then(r => {
		if (r.ok) {
			return r.json();
		} else {
			throw new Error(`${r.status} - ${r.statusText} (${file})`);
		}
	});
}

function load() {
	location.href = `${location.origin}?year=${year}&week=${week}`;
}

function changeWeek(func) {
	const d = func(moment(`${year} ${week}`, "GGGG WW"));
	year = d.format("GGGG");
	week = d.format("WW");
	load();
}

function currWeek() {
	changeWeek(() => moment());
}

function prevWeek() {
	changeWeek(d => d.subtract(1, "week"));
}

function nextWeek() {
	changeWeek(d => d.add(1, "week"));
}

function format(d) {
	return d.format("HH:mm:ss");
}

function renderData(data) {

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

function renderError(e) {
	errorElement.innerHTML = e;
}

(function() {

	const url = new URL(location.href);
	year = url.searchParams.get("year");
	week = url.searchParams.get("week");

	if (!(year && week)) {
		currWeek();
		return;
	}

	loadFile().then(renderData).catch(renderError);
	weekElement.innerHTML = `${year}-W${week}`;

})();

