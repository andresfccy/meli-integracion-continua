import {defineConfig, devices} from '@playwright/test';

export default defineConfig({
  testDir: './e2e',
  fullyParallel: false,
  workers: 1,
  reporter: 'list',
  use: {
    baseURL: 'http://127.0.0.1:3000',
    trace: 'retain-on-failure',
  },
  webServer: [
    {
      command: 'go run ./cmd/api',
      cwd: '../backend',
      url: 'http://127.0.0.1:8080/api/health',
      reuseExistingServer: true,
      timeout: 120_000,
    },
    {
      command: 'pnpm run dev',
      cwd: '.',
      url: 'http://127.0.0.1:3000/login',
      reuseExistingServer: true,
      timeout: 120_000,
    },
  ],
  projects: [{name: 'chromium', use: {...devices['Desktop Chrome']}}],
});
