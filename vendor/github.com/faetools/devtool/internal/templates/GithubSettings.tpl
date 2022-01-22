# Code generated by devtool; DO NOT EDIT.
repository:
  topics: {{ if .UsesGo -}} golang, {{ end -}} devtool
  private: true
  has_wiki: false
  has_downloads: false
  allow_merge_commit: false
  delete_branch_on_merge: true
teams:
  - name: engineering
    permission: maintain
  - name: {{ .Team }}
    permission: maintain
{{- if .IsProductionTraffic }}
branches:
  - name: master
    protection:
      required_pull_request_reviews:
        required_approving_review_count: {{ if eq .Tier "Tier 1 - Critical" }}2{{else}}1{{end}}
        dismiss_stale_reviews: true
        require_code_owner_reviews: true
      required_status_checks:
        strict: true
        contexts:
					- "SonarCloud Code Analysis"
          # - "ci/circleci: Go Kit Validate"
          - "ci/circleci: Unit Tests"
          - "ci/circleci: Contract Tests"
          - "ci/circleci: Service Compiles"
        enforce_admins: false
      restrictions: null
  - name: main
    protection:
      required_pull_request_reviews:
        required_approving_review_count: {{if eq .Tier "Tier 1 - Critical"}}2{{else}}1{{end}}
        dismiss_stale_reviews: true
        require_code_owner_reviews: true
      required_status_checks:
        strict: true
        contexts:
					- "SonarCloud Code Analysis"
          # - "ci/circleci: Go Kit Validate"
          - "ci/circleci: Unit Tests"
          - "ci/circleci: Contract Tests"
          - "ci/circleci: Service Compiles"
      enforce_admins: false
      restrictions: null
{{- else }}
branches:
  - name: master
    protection:
      required_pull_request_reviews:
				dismiss_stale_reviews: true
        require_code_owner_reviews: true
      #  required_approving_review_count: 0
        enforce_admins: false
      {{ if .UsesGo -}}
			required_status_checks:
        strict: true
        contexts:
          # - 'ci/circleci: Go Kit Validate'
          # - "SonarCloud Code Analysis"
          - "ci/circleci: Unit Tests"
          - "ci/circleci: Contract Tests"
					{{- if and false .IsMicroservice }}
          - "ci/circleci: Service Compiles"
					{{- end }}
        enforce_admins: false
			{{- end }}
      restrictions: null
  - name: main
    protection:
      required_pull_request_reviews:
        dismiss_stale_reviews: true
        require_code_owner_reviews: true
        # required_approving_review_count: 0
        enforce_admins: false
      {{ if .UsesGo -}}
      required_status_checks:
        strict: true
        contexts:
          # - 'ci/circleci: Go Kit Validate'
          # - "SonarCloud Code Analysis"
          - "ci/circleci: Unit Tests"
          - "ci/circleci: Contract Tests"
					{{- if and false .IsMicroservice }}
          - "ci/circleci: Service Compiles"
					{{- end }}
        enforce_admins: true
			{{- end }}
      restrictions: null
{{- end }}
