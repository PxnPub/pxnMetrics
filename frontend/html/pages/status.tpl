
<!-- Broker Offline -->
<div id="broker-offline-alert" class="alert alert-danger alert-dismissible" style="display: none;" role="alert">
	<h2>
		<i class="bi bi-exclamation-triangle-fill"></i>
		Connection to broker has failed
		<!--&nbsp;<span class="fs-6">(2 minutes)</span>-->
	</h2>
</div>

<!-- Broker Warning -->
<div id="broker-offline-warning" class="alert alert-warning alert-dismissible fade show" style="display: none;" role="alert">
	<h5><i class="bi bi-exclamation-triangle-fill"></i>
		Connection to broker failed
		<!--&nbsp;<span class="fs-6">(3 minutes ago)</span>-->
	</h5>
	<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>

<!-- Broker Uptime Bar -- >
<div class="progress" style="width: 80%; height: 20px; margin-top: 20px;
margin-bottom: 30px; margin-left: auto; margin-right: auto;">
	<div class="progress-bar bg-success" style="width: 55%"></div>
	<div class="progress-bar bg-success" style="width: 5%"></div>
	<div class="progress-bar bg-warning progress-bar-striped" style="width: 5%"
		data-bs-toggle="tooltip" data-bs-placement="top" data-bs-title="5 minutes"></div>
	<div class="progress-bar bg-success" style="width: 14%"></div>
	<div class="progress-bar bg-danger progress-bar-striped"  style="width: 8%"
		data-bs-toggle="tooltip" data-bs-placement="top" data-bs-title="10 minutes"></div>
	<div class="progress-bar bg-success" style="width: 18%"></div>
</div>-->

<!-- Shards Container -->
<div class="container-fluid mt-3">
	<div class="row justify-content-center" id="shard-status-container">

		<!-- Shard Status -->
		<div class="col-1 pt-2 pb-2 text-center shard-status-template" style="display: none;">
			<!-- Name -->
			<h3 class="shard-name">&lt;Name&gt;</h3>
			<!-- Status -->
			<h3><span class="badge shard-status-button">&lt;Status&gt;</span></h3>
			<!-- Uptime Bar -- >
			<div class="progress" style="width: 100%; height: 20px;">
				<div class="progress-bar bg-success" style="width: 100%"></div>
			</div>-->
			<!-- Request Stats -->
			<table>
				<tr><td align="right">Last&nbsp;Batch:</td><td width="8px"></td>
					<td align="left" class="shard-last-batch fs-5">&lt;Last-Batch&gt;</td></tr>
				<tr><td align="right" class="shard-requests">Requests:</td><td></td>
					<td align="left" class="shard-queue-waiting fs-5">&lt;Queued&gt;</td></tr>
				<tr><td></td><td></td>
					<td align="left" class="shard-req-sec-min fs-6">&lt;Req-Per&gt;</td></tr>
				<tr><td colspan="3" class="shard-req-total fs-6">&lt;Total-Req&gt;</td></tr>
			</table>
		</div>

<!--
<div class="progress" style="width: 100%; height: 20px;">

	<div class="progress-bar bg-success" style="width: 30%"></div>

	<div class="progress-bar bg-danger progress-bar-striped"  style="width: 8%"
		data-bs-toggle="tooltip" data-bs-placement="top" data-bs-title="10 minutes"></div>

	<div class="progress-bar bg-success" style="width: 35%"></div>

	<div class="progress-bar bg-warning progress-bar-striped" style="width: 15%"
		data-bs-toggle="tooltip" data-bs-placement="top" data-bs-title="5 minutes"></div>

	<div class="progress-bar bg-success" style="width: 25%"></div>

	<div class="progress-bar bg-warning progress-bar-striped" style="width: 20%"
		data-bs-toggle="tooltip" data-bs-placement="top" data-bs-title="5 minutes"></div>

</div>
-->

	</div>
</div>
