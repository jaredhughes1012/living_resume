import { expect, test, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom'
import CreateIdentityForm from './CreateIdentityForm';
import userEvent from '@testing-library/user-event'
import { IdentityInput } from '@types';

test('CreateIdentityForm submit on button press', async () => {
  const user = userEvent.setup();
  const expected: IdentityInput = {
    activationCode: '123ABC',
    firstName: 'John',
    lastName: 'Doe',
    accountId: '123',
    credentials: {
      email: 'test@test.com',
      password: 'password',
    }
  };

  const onContinue = vi.fn();
  render(
    <CreateIdentityForm
      email={expected.credentials.email}
      code={expected.activationCode}
      onChange={vi.fn()}
      onSubmit={onContinue} />
  );

  const firstNameInput = await screen.findByLabelText('First Name');
  const lastNameInput = await screen.findByLabelText('Last Name');
  const accountIdInput = await screen.findByLabelText('Account ID');
  const passwordInput = await screen.findByLabelText('Password');

  await user.type(firstNameInput, expected.firstName);
  await user.type(lastNameInput, expected.lastName);
  await user.type(accountIdInput, expected.accountId);
  await user.type(passwordInput, expected.credentials.password);


  await user.click(screen.getByText('Create Account'));
  expect(onContinue).toHaveBeenCalledWith(expected);
});

test('CreateIdentityForm submit on enter', async () => {
  const user = userEvent.setup();
  const expected: IdentityInput = {
    activationCode: '123ABC',
    firstName: 'John',
    lastName: 'Doe',
    accountId: '123',
    credentials: {
      email: 'test@test.com',
      password: 'password',
    }
  };

  const onContinue = vi.fn();
  render(
    <CreateIdentityForm
      email={expected.credentials.email}
      code={expected.activationCode}
      onChange={vi.fn()}
      onSubmit={onContinue} />
  );

  const firstNameInput = await screen.findByLabelText('First Name');
  const lastNameInput = await screen.findByLabelText('Last Name');
  const accountIdInput = await screen.findByLabelText('Account ID');
  const passwordInput = await screen.findByLabelText('Password');

  await user.type(firstNameInput, expected.firstName);
  await user.type(lastNameInput, expected.lastName);
  await user.type(accountIdInput, expected.accountId);
  await user.type(passwordInput, expected.credentials.password);
  await user.type(passwordInput, '{enter}');

  expect(onContinue).toHaveBeenCalledWith(expected);
});