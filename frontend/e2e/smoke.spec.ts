import {expect, Page, test} from '@playwright/test';

async function login(page: Page, role: 'Vendedor' | 'Comprador' | 'Operaciones') {
  await page.goto('/login');
  await page.getByRole('button', {name: role}).click();
  await page.getByRole('button', {name: 'Continuar'}).click();
}

async function logout(page: Page) {
  await page.getByRole('button', {name: 'Cerrar sesión'}).click();
}

test.describe.serial('documented role flows', () => {
  test('rejects invalid credentials and logs in a seeded seller', async ({page}) => {
    await page.goto('/login');
    await page.getByLabel('Correo').fill('unknown@example.com');
    await page.getByLabel('Contraseña').fill('incorrect');
    await page.getByRole('button', {name: 'Continuar'}).click();
    await expect(page.locator('.error')).toBeVisible();

    await page.getByRole('button', {name: 'Vendedor'}).click();
    await page.getByRole('button', {name: 'Continuar'}).click();
    await expect(page).toHaveURL(/\/seller\/orders$/);
    await expect(page.getByRole('heading', {name: 'Órdenes recientes'})).toBeVisible();
  });

  test('buyer asks, seller answers, and buyer approves without page reloads', async ({page}) => {
    await login(page, 'Comprador');
    await page.getByRole('link', {name: /Monitor 27 pulgadas/}).click();
    await page.getByPlaceholder('Escribe tu pregunta…').fill('¿Puedo confirmar el estado del monitor?');
    await page.getByRole('button', {name: 'Enviar pregunta'}).click();
    await expect(page.getByText('Pregunta enviada.')).toBeVisible();
    await logout(page);

    await login(page, 'Vendedor');
    await page.getByRole('link', {name: /Monitor 27 pulgadas/}).click();
    await page.getByLabel('Respuesta').fill('Sí, el monitor fue entregado correctamente.');
    await page.getByRole('button', {name: 'Responder'}).click();
    await expect(page.getByText('Respuesta guardada.')).toBeVisible();
    await logout(page);

    await login(page, 'Comprador');
    await page.getByRole('link', {name: /Monitor 27 pulgadas/}).click();
    await page.getByRole('button', {name: 'Aprobar respuesta'}).click();
    await expect(page.getByText('Respuesta aprobada y pregunta cerrada.')).toBeVisible();
  });

  test('Operations filters sellers by name, email and id and runs a manual review', async ({page}) => {
    await login(page, 'Operaciones');
    for (const value of ['Mercado Norte', 'seller@example.com', 'seller-1']) {
      await page.getByLabel('Vendedor').fill(value);
      await page.getByRole('button', {name: 'Aplicar filtros'}).click();
      await expect(page.getByText(/Mercado Norte \(seller-1\)/)).toBeVisible();
    }
    await page.getByRole('button', {name: 'Ejecutar revisión'}).click();
    await expect(page.locator('.success')).toContainText('completed');
    await expect(page.locator('.success')).toContainText(/lotes enviados/);
  });
});
