package site

const (
	nginxConf = `
{{if .Ssl}}
	server {
		listen 80;
		server_name {{.Name}};
		rewrite ^(.*) https://$host$1 permanent;
	}
{{end}}
	upstream {{.Name}}_prod {
		server localhost:{{.Port}} fail_timeout=0;
	}
	server {
{{if .Ssl}}
		listen 443;
		ssl  on;
		ssl_certificate  /etc/ssl/certs/{{.Name}}.crt;
		ssl_certificate_key  /etc/ssl/private/{{.Name}}.key;
		ssl_session_timeout  5m;
		ssl_protocols  SSLv2 SSLv3 TLSv1;
		ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
		ssl_prefer_server_ciphers  on;
{{else}}
		listen 80;
{{end}}
		client_max_body_size 4G;
		keepalive_timeout 10;
		proxy_buffers 16 64k;
		proxy_buffer_size 128k;
		server_name {{.Name}};
		root {{.Root}};
		index index.html;
		access_log /var/log/nginx/{{.Name}}.access.log;
		error_log /var/log/nginx/{{.Name}}.error.log;
		try_files $uri/index.html $uri @{{.Name}}_prod;

		location ^~ /(assets|attachments)/ {
			gzip_static on;
			expires max;
			access_log off;
			add_header Cache-Control "public";
		}

		location ~* \.(?:rss|atom)$ {
			expires 12h;
			access_log off;
			add_header Cache-Control "public";
		}

		location @{{.Name}}_prod {
		{{if .Ssl}}
			proxy_set_header X-Forwarded-Proto https;
		{{else}}
			proxy_set_header X-Forwarded-Proto http;
		{{end}}
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_redirect off;
			proxy_pass http://{{.Name}}_prod;
			# limit_req zone=one;
		}

		if ($request_method !~ ^(GET|HEAD|PUT|PATCH|POST|DELETE|OPTIONS)$ ){
			return 405;
		}
		if (-f $document_root/system/maintenance.html) {
			return 503;
		}
	}
	`

	robotsTxt = `
User-agent: *
Disallow:
Crawl-delay: 10
Sitemap: {{.Home}}/sitemap.xml.gz
`
)
