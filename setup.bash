#!/bin/bash

# Root directories
mkdir -p \
apps \
packages \
infra \
deployments \
scripts \
tools \
docs \
.github/workflows

# Applications
mkdir -p \
apps/api/cmd/server \
apps/api/internal/{handlers,middleware,services,dto,auth,cache,config} \
apps/batch/cmd/batch \
apps/batch/internal/{ingestion,indicators,intelligence,opportunities,insights,scheduler} \
apps/worker \
apps/mobile

# Domain Packages
mkdir -p \
packages/master \
packages/marketdata \
packages/indicators \
packages/fundamentals \
packages/configuration \
packages/intelligence \
packages/opportunities \
packages/insights \
packages/users \
packages/events

# Shared Package
mkdir -p \
packages/shared/{logger,errors,config,timeutil,uuid,constants}

# Infrastructure
mkdir -p \
infra/docker \
infra/db/{migrations,seeds,schema} \
infra/monitoring \
infra/redis

# Deployments
mkdir -p \
deployments/{dev,staging,prod}

# Scripts
mkdir -p \
scripts/{seed-data,backfill,migration,maintenance}

# Tools
mkdir -p \
tools/{generators,lint,format}

# Documentation
mkdir -p \
docs/{01-vision,02-hld,03-erd,04-pdd,05-lld,06-api,07-sprints}

# Migration placeholders
touch \
infra/db/migrations/000001_master_domain.sql \
infra/db/migrations/000002_market_data.sql \
infra/db/migrations/000003_indicator_domain.sql \
infra/db/migrations/000004_fundamental_domain.sql \
infra/db/migrations/000005_configuration_domain.sql \
infra/db/migrations/000006_intelligence_domain.sql \
infra/db/migrations/000007_user_domain.sql \
infra/db/migrations/000008_insights_domain.sql

# Event contracts placeholders
touch \
packages/events/market_data_ingested.go \
packages/events/indicators_calculated.go \
packages/events/market_regime_calculated.go \
packages/events/stock_scores_calculated.go \
packages/events/opportunity_detected.go \
packages/events/insights_generated.go

# Root files
touch \
README.md \
Makefile \
docker-compose.yml \
.env.example \
go.work

# GitHub Actions
touch \
.github/workflows/ci.yml

echo "✅ MarketPilot repository structure created successfully."