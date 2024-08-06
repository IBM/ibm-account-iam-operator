package yamls

var IMConfigYamls = []string{
	IM_CONFIG_ROLE,
	IM_CONFIG_ROLE_BINDING,
	IM_CONFIG_SA,
	IM_CONFIG_JOB,
}

var IM_CONFIG_JOB = `
apiVersion: batch/v1
kind: Job
metadata:
  name: mcsp-im-config-job
  labels:
    app: mcsp-im-config-job
spec:
  template:
    metadata:
      labels:
        app: mcsp-im-config-job
    spec:
      containers:
      - name: mcsp-im-config-job
        image: MCSP_IM_CONFIG_JOB_IMAGE
        command: ["./mcsp-im-config-job"]
        imagePullPolicy: Always
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
              - ALL
          runAsNonRoot: true
          seccompProfile:
            type: RuntimeDefault
        env:
          - name: LOG_LEVEL
            value: debug
          - name: NAMESPACE
            value: {{ .AccountIAMNamespace }}
          - name: IM_HOST_BASE_URL
            value: {{ .IAMHostURL }}
          - name: ACCOUNT_IAM_BASE_URL
            value: {{ .AccountIAMURL }}
          - name: ACCOUNT_IAM_CONSOLE_BASE_URL
            value: {{ .AccountIAMHostURL }}
      serviceAccountName: mcsp-im-config-sa
      restartPolicy: OnFailure
`

var IM_CONFIG_SA = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: mcsp-im-config-sa
  labels:
    app: mcsp-im-config-sa
`

var IM_CONFIG_ROLE = `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: mcsp-im-config-role
  labels:
    app: mcsp-im-config-role
rules:
  - apiGroups: [""]
    resources: ["pods", "configmaps", "secrets"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["create", "update", "delete"]
`

var IM_CONFIG_ROLE_BINDING = `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: mcsp-im-config-rb
  labels:
    app: mcsp-im-config-rb
subjects:
  - kind: ServiceAccount
    name: mcsp-im-config-sa
roleRef:
  kind: Role
  name: mcsp-im-config-role
  apiGroup: rbac.authorization.k8s.io
`
