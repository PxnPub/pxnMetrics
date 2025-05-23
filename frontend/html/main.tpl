<nav class="nav navbar navbar-expand-sm bg-gradient">
	<div class="container-fluid">
		<a class="navbar-brand" href="#">
			<img src="/static/metrics-logo.png" />
			pxnMetrics
		</a>
		<button class="navbar-toggler" type="button"
				data-bs-toggle="collapse"
				data-bs-target="#NavbarSupportedContent"
				aria-controls="NavbarSupportedContent"
				aria-expanded="false"
				aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="NavbarSupportedContent">
			<ul class="navbar-nav me-auto mb-2 mb-sm-0">
				<li class="nav-item">
					<button href="#" class="nav-link{{if eq .Page "home"}} active{{end}}">Home</button>
				</li>
				<li class="nav-item">
					<button href="#" class="nav-link{{if eq .Page "metrics"}} active{{end}}">Metrics</button>
				</li>
				<li class="nav-item">
					<button href="#" class="nav-link{{if eq .Page "wiki"}} active{{end}}">Wiki/Docs</button>
				</li>
				<li class="nav-item">
					<button href="#" class="nav-link{{if eq .Page "blog"}} active{{end}}">News/Blog</button>
				</li>
				<li class="nav-item">
					<button href="#" class="nav-link{{if eq .Page "about"}} active{{end}}">About Us</button>
				</li>
			</ul>
		</div>
	</div>
</nav>
