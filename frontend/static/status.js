
const shard_status_container = document.getElementById('shard-status-container');
const shard_status_template  = document.querySelector('.shard-status-template');
var offline_alert_timer = 1;

function FetchAndUpdateStatus() {
	if (!shard_status_container) console.error('class not found: .shard-status-container');
	if (!shard_status_template)  console.error('class not found: .shard-status-template');
	fetch('/api/status/', { signal: AbortSignal.timeout(5000) })
		.then(resp => {
			if (!resp.ok) { throw new Error('Fetch result not ok'); }
			return resp.json();
		})
		.then(json => {
			UpdateStatus(json);
		})
		.catch(err => {
			SetStatusOffline();
			console.log('Fetch error: ', err)
		});
}

function SetStatusOffline() {
	if (offline_alert_timer !== null) {
		clearInterval(offline_alert_timer);
		offline_alert_timer = null;
	}
	document.getElementById('broker-offline-warning').style.display = 'none';
	document.getElementById('broker-offline-alert'  ).style.display = 'block';
}

function ResetStatusOffline() {
	document.getElementById('broker-offline-warning').style.display = 'none';
	document.getElementById('broker-offline-alert'  ).style.display = 'none';
}

function UpdateStatus(json) {
	if (offline_alert_timer === null) {
		offline_alert_timer = setInterval(ResetStatusOffline, 10000);
		document.getElementById('broker-offline-warning').style.display = 'block';
		document.getElementById('broker-offline-alert'  ).style.display = 'none';
	}
	const frags = document.createDocumentFragment();
	json.Shards.forEach(entry => {
		const clone = shard_status_template.cloneNode(true);
		clone.style.display = 'block';
		clone.querySelector('.shard-name').textContent = entry.Name;
		UpdateStatusButton   (entry, clone.querySelector('.shard-status-button'));
		UpdateStatusLastBatch(entry,
			clone.querySelector('.shard-last-batch-title'),
			clone.querySelector('.shard-last-batch-value'));
		UpdateStatusWaiting  (entry, clone.querySelector('.shard-requests'),
			clone.querySelector('.shard-queue-waiting'));
		UpdateStatusPerPeriod(entry, clone.querySelector('.shard-req-sec-min'  ));
		UpdateStatusTotalReq (entry, clone.querySelector('.shard-req-total'    ));
		frags.appendChild(clone);
	});
	shard_status_container.replaceChildren(frags);
}

function UpdateStatusButton(entry, field) {
	switch (entry.Status) {
	case 'Online':
		const now = new Date();
		field.innerHTML = `<i class="bi bi-emoji-sunglasses"></i>` +
			`&nbsp;Online&nbsp;`+(now.getMonth()===3 && now.getDate()===1
			? shard_status_online_1 : `<i class="bi bi-hand-thumbs-up"></i>`);
		field.classList.add('text-bg-success');
		break;
	case 'Alert':
		field.innerHTML = `<i class="bi bi-emoji-grimace"></i>` +
			`&nbsp;Online&nbsp;<i class="bi bi-question-circle"></i>`;
		field.classList.add('text-bg-warning');
		break;
	case 'Offline':
		field.innerHTML = `<i class="bi bi-exclamation-triangle-fill"></i>` +
			`&nbsp;Offline&nbsp;<i class="bi bi-exclamation-triangle-fill"></i>`;
		field.classList.add('text-bg-danger');
		break;
	default:
		field.innerHTML = status;
		break;
	}
}

function UpdateStatusLastBatch(entry, title, field) {
	if (entry.Status === 'Offline' && entry.LastBatch === 0) {
		title.style.display = 'none';
		field.textContent = '';
	} else {
		title.style.display = 'block';
		field.innerHTML = FormatUptimeSeconds(entry.LastBatch) + (
			(entry.BatchWaiting === 0 && entry.Status === 'Offline') ? ''
			: ` <font size="-1">(` + entry.BatchWaiting + `&nbsp;waiting)</font>`
		);
	}
}

function UpdateStatusWaiting(entry, fieldA, fieldB) {
	if (entry.Status === 'Offline') {
		fieldA.textContent = '';
		fieldB.textContent = '';
	} else {
		fieldA.textContent = 'Requests';
		fieldB.textContent = entry.QueueWaiting + ' Queued';
	}
}

function UpdateStatusPerPeriod(entry, field) {
	if (entry.Status === 'Offline') {
		field.textContent = '';
	} else {
		field.innerHTML =
			entry.ReqPerSec + ` <font size="-1">/sec</font><br />` +
			entry.ReqPerMin + ` <font size="-1">/min</font><br />` +
			entry.ReqPerHour+ ` <font size="-1">/hr</font><br />` +
			entry.ReqPerDay + ` <font size="-1">/day</font>`;
	}
}

function UpdateStatusTotalReq(entry, field) {
	if (entry.Status === 'Offline') {
		field.textContent = '';
	} else {
		field.innerHTML = '<font size="-1">Total Req: </font>' +
			entry.ReqTotal.toLocaleString('en-US');
	}
}

function FormatUptimeSeconds(seconds) {
	const d = Math.floor( seconds / 86400     );
	const h = Math.floor( seconds / 3600      );
	const m = Math.floor((seconds % 3600) / 60);
	const s = seconds % 60;
	return '' +
		(d > 0 ? String(d)+'d '                 :        '') +
		(h > 0 ? String(h).padStart(2, '0')+':' :        '') +
		(h > 0 ? String(m).padStart(2, '0')     : String(m)) +
		':'+String(s).padStart(2, '0');
}



FetchAndUpdateStatus();
setInterval(FetchAndUpdateStatus, 1000);

const shard_status_online_1 = `<svg width="24" height="24" fill="white" viewBox="0 0 32 32"><path ` +
	`d="M 16 2 C 14.355469 2 13 3.355469 13 5 L 13 10.1875 C 12.683594 10.074219 12.351563 10 12 `  +
	`10 C 10.355469 10 9 11.355469 9 13 L 9 16.65625 L 6.90625 19.34375 C 5.628906 20.996094 `      +
	`5.714844 23.367188 7.09375 24.9375 L 9.46875 27.625 C 10.796875 29.136719 12.707031 30 `       +
	`14.71875 30 L 20 30 C 23.855469 30 27 26.855469 27 23 L 27 14 C 27 12.355469 25.644531 11 24 ` +
	`11 C 23.464844 11 22.96875 11.15625 22.53125 11.40625 C 21.996094 10.5625 21.0625 10 20 10 C ` +
	`19.648438 10 19.316406 10.074219 19 10.1875 L 19 5 C 19 3.355469 17.644531 2 16 2 Z M 16 4 C ` +
	`16.566406 4 17 4.433594 17 5 L 17 15 L 19 15 L 19 13 C 19 12.433594 19.433594 12 20 12 C `     +
	`20.566406 12 21 12.433594 21 13 L 21 15 L 23 15 L 23 14 C 23 13.433594 23.433594 13 24 13 C `  +
	`24.566406 13 25 13.433594 25 14 L 25 23 C 25 25.773438 22.773438 28 20 28 L 14.71875 28 C `    +
	`13.28125 28 11.917969 27.394531 10.96875 26.3125 L 8.59375 23.59375 C 7.839844 22.734375 `     +
	`7.800781 21.5 8.5 20.59375 L 9 19.9375 L 9 21 L 11 21 L 11 13 C 11 12.433594 11.433594 12 12 ` +
	`12 C 12.566406 12 13 12.433594 13 13 L 13 15 L 15 15 L 15 5 C 15 4.433594 15.433594 4 16 4 Z`  +
	`"></svg>`;
