const errorElement = document.getElementById("error");
const dataElement = document.getElementById("data");

const dataDir = "data";

let year = 0;
let week = 0;

async function loadFile() {

	let w = week;
	if (w < 10) {
		w = `0${w}`;
	}

	const file = `${dataDir}/${year}/${w}.json`;

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
	year = parseInt(d.format("GGGG"));
	week = parseInt(d.format("W"));
	load();
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

		dataElement.innerText += `${date}\n`;

		if (day.Activities && day.Activities.length > 0) {
			let d = 0;
			dataElement.innerText += "\n";
			for (const act of day.Activities) {
				const start = moment(act.Start);
				const end = moment(act.End);
				const dur = end.diff(start);
				d += dur;
				dataElement.innerText += `\t${act.Title}:\t\t${format(start)} - ${format(end)}\t${format(moment.utc(dur))}\n`;
			}

			const dur = moment.utc(moment.duration(d).asMilliseconds());
			dataElement.innerText += `\n\t\t\t\t\t\t${format(dur)}\n`;
		}

		let d = 0;
		dataElement.innerText += "\n";
		for (const span of day.Spans) {
			const start = moment(span.Start);
			const end = moment(span.End);
			const dur = end.diff(start);
			d += dur;
			dataElement.innerText += `\t${format(start)} - ${format(end)}\t${format(moment.utc(dur))}\n`;
		}
		const dur = moment.utc(moment.duration(d).asMilliseconds());
		dataElement.innerText += `\n\t\t\t\t${format(dur)}\n\n`;
	}
}

function renderError(e) {
	errorElement.innerText = e;
}

(function() {

	const url = new URL(location.href);
	year = parseInt(url.searchParams.get("year"));
	week = parseInt(url.searchParams.get("week"));

	if (!(year && week)) {
		const d = moment();
		year = parseInt(d.format("GGGG"));
		week = parseInt(d.format("W"));
		load(year, week);
		return;
	}

	loadFile().then(renderData).catch(renderError);

})();

