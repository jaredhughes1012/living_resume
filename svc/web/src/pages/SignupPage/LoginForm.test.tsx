import { expect, test, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import userEvent from '@testing-library/user-event';
import { Credentials } from '@types';
import LoginForm from './LoginForm';

test('CreateIdentityForm submit on button press', async () => {
  const user = userEvent.setup();
  const expected: Credentials = {
    email: 'test@test.com',
    password: 'password',
  };

  const onContinue = vi.fn();
  render(
    <LoginForm onSubmit={onContinue} onChange={vi.fn()} />
  );

  const emailInput = await screen.findByLabelText('Email');
  const passwordInput = await screen.findByLabelText('Password');

  await user.type(emailInput, expected.email);
  await user.type(passwordInput, expected.password);

  await user.click(screen.getByText('Log In'));
  expect(onContinue).toHaveBeenCalledWith(expected);
});

test('CreateIdentityForm submit on enter', async () => {
  const user = userEvent.setup();
  const expected: Credentials = {
    email: 'test@test.com',
    password: 'password',
  };

  const onContinue = vi.fn();
  render(
    <LoginForm onSubmit={onContinue} onChange={vi.fn()} />
  );

  const emailInput = await screen.findByLabelText('Email');
  const passwordInput = await screen.findByLabelText('Password');

  await user.type(emailInput, expected.email);
  await user.type(passwordInput, expected.password);
  await user.type(passwordInput, '{enter}');

  await user.click(screen.getByText('Log In'));
  expect(onContinue).toHaveBeenCalledWith(expected);
});