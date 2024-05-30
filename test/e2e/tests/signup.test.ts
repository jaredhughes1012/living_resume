import { expect, test } from '@playwright/test';

test('User can log in', async ({ page }) => {
  const runId = process.env.RUN_ID!;
  const input = {
    email: `${runId}@test.com`,
    password: 'P4$$w0rd!1',
  }

  await page.goto('/jared-hughes');
  await page.getByRole('button', { name: 'Log in' }).click();

  const p = page.waitForResponse(resp => resp.url().includes('authenticate') && resp.status() === 200);
  await page.getByLabel('Email').fill(input.email);
  await page.getByLabel('Password').fill(input.password);
  await page.getByRole('button', { name: "Log In" }).click();

  const res = await p;
  const response = await res.json();

  expect(response.token).toBeTruthy();
  expect(response.identity).toBeTruthy();
});