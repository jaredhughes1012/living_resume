import { expect, test } from 'vitest';
import { render, screen } from '@testing-library/react';
import OrDivider from './OrDivider';
import '@testing-library/jest-dom'

test('OrDivider displays the text "or"', async () => {
  render(<OrDivider />);

  const input = await screen.findByText('OR');
  expect(input).toBeInTheDocument();
});