#!/bin/bash
set -euo pipefail

wire gen ./pkg/config
wire gen ./pkg/tools
wire gen ./pkg/tools/az
wire gen ./pkg/provisioning
wire gen ./pkg/provisioning/bicep
wire gen ./pkg/provisioning/terraform
wire gen ./pkg/actions
