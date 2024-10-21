import { render, screen } from '@testing-library/react';
import { App } from './App';

test('should render App', () => {
  render(<App />);
  expect(screen.getByText(/Sign in to access your notifications/i)).toBeInTheDocument();
});
