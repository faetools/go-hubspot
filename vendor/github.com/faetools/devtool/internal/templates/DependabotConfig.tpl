# Code generated by devtool; DO NOT EDIT.
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
      day: "sunday"
      time: "{{ .DependabotScheduledAt }}"
    reviewers:
      - "{{ .Team }}"
    ignore:
     # - dependency-name: "github.com/faetools/xxx"
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "daily"
    reviewers:
      - "{{ .Team }}"
