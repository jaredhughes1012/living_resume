import { test as setup, expect } from '@playwright/test';

const authFile = 'playwright/.auth/user.json';

setup('authenticate', async ({ page }) => {
  const runId = process.env.RUN_ID || new Date().toISOString().replace(/[^0-9]/g, '');
  process.env.RUN_ID = runId;

  const input = {
    accountId: runId,
    firstName: 'John',
    lastName: 'Doe',
    credentials: {
      email: `${runId}@test.com`,
      password: 'P4$$w0rd!1',
    }
  };

  await page.goto('/signup');

  const p = page.waitForResponse(resp => resp.url().includes('accounts/initiate') && resp.status() === 200);
  await page.getByLabel('Email').fill(input.credentials.email);
  await page.getByText('Continue').click();

  const res = await p;
  const debugData = await res.json();
  expect(res.status()).toBe(200);

  await page.goto(debugData.url);
  await page.getByLabel('First name').fill(input.firstName);
  await page.getByLabel('Last name').fill(input.lastName);
  await page.getByLabel('Password').fill(input.credentials.password);
  await page.getByLabel('Account ID').fill(input.accountId);
  await page.getByText('Create Account').click();

  await page.context().storageState({ path: authFile });
});