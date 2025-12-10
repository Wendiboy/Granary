
        function renderTable(jsonText) {
            let spends;
            try { spends = JSON.parse(jsonText); }
            catch (e) {
                document.getElementById('spends-table').innerHTML = '<p style="color:red">Ошибка загрузки данных</p>';
                return;
            }

            if (!spends.length) {
                document.getElementById('spends-table').innerHTML = '<p class="empty">Пока нет расходов<br>Добавьте первый!</p>';
                return;
            }

            spends.sort((a, b) => new Date(b.date) - new Date(a.date));

            const currencySymbol = cur => ({ RUB: '₽', USD: '$', BTC: '₿', ETH: 'Ξ' }[cur] || cur);
            const formatDate = d => new Date(d).toISOString().split('T')[0];

            const rows = spends.map(s => `
            <tr id="row-${s.id}">
                <td data-label="Дата">${new Date(s.date).toLocaleDateString('ru-RU')}</td>
                <td data-label="Категория">${s.category || '—'}</td>
                <td data-label="Счёт">${s.account || '—'}</td>
                <td data-label="Описание" class="note">${s.note || '—'}</td>
                <td data-label="Тег">${s.labels ? `<span class="tag">${s.labels}</span>` : '—'}</td>
                <td class="amount">${parseFloat(s.amount).toLocaleString('ru-RU')} ${currencySymbol(s.currency)}</td>
                <td class="actions" style="white-space: nowrap; text-align: center;">
                    <button class="edit-btn" title="Редактировать"
                            onclick="editRow('${s.id}', '${formatDate(s.date)}', '${(s.category || '').replace(/'/g, "\\'")}', '${(s.account || '').replace(/'/g, "\\'")}', '${s.amount}', '${s.currency || 'RUB'}', '${(s.labels || '').replace(/'/g, "\\'")}', '${(s.note || '').replace(/'/g, "&#39;")}')">
                        Edit
                    </button>
                    <button class="delete-btn" title="Удалить"
                            hx-delete="/spends/${s.id}"
                            hx-confirm="Точно удалить этот расход?"
                            hx-target="#row-${s.id}"
                            hx-swap="outerHTML"
                            hx-on::after-request="if(event.detail.successful) htmx.trigger('#spends-table', 'reload')">
                        Delete
                    </button>
                </td>
            </tr>
        `).join('');

            document.getElementById('spends-table').innerHTML = `
            <table>
                <thead>
                    <tr>
                        <th>Дата</th><th>Категория</th><th>Счёт</th><th>Описание</th><th>Тег</th><th>Сумма</th>
                        <th style="width:110px; text-align:center;">Действия</th>
                    </tr>
                </thead>
                <tbody>${rows}</tbody>
            </table>
        `;

            htmx.process(document.getElementById('spends-table'));
        }

        function editRow(id, date, category, account, amount, currency, labels, note) {
    const row = document.getElementById('row-' + id);

    row.innerHTML = `
    <td colspan="7" style="padding:8px; background:#f9f9f9;" >
        <form hx-patch="/spends/${id}"
              hx-ext="json-enc"
              hx-target="#spends-table"
              hx-swap="innerHTML"
              style="display:grid; grid-template-columns:repeat(auto-fit,minmax(140px,1fr)); gap:8px; align-items:end;">
            
            <input type="date" name="date" value="${date}" required>

            <select name="category_id" required>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"          ${category === 'Еда' ? 'selected' : ''}>Еда</option>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"    ${category === 'Транспорт' ? 'selected' : ''}>Транспорт</option>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"  ${category === 'Развлечения' ? 'selected' : ''}>Развлечения</option>
            </select>

            <select name="account_id">
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"        ${account === 't-bank' ? 'selected' : ''}>T-bank</option>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"          ${account === 'sber' ? 'selected' : ''}>SBER</option>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"     ${account === 'alfa_visa' ? 'selected' : ''}>Alfa Visa</option>
                <option value="87f5012e-6a4d-4d55-8fa4-7705598fb85c"      ${account === 'alfa_mir' ? 'selected' : ''}>Alfa MIR</option>
            </select>

            <input type="text" name="note" value="${note}" placeholder="Описание">
            <input type="text" name="labels" value="${labels}" placeholder="Тег">
            <input type="number" name="amount" step="0.01" value="${amount}" required>

            <select name="currency">
                <option value="RUB" ${currency === 'RUB' ? 'selected' : ''}>₽ RUB</option>
                <option value="USD" ${currency === 'USD' ? 'selected' : ''}>$ USD</option>
                <option value="BTC" ${currency === 'BTC' ? 'selected' : ''}>₿ BTC</option>
                <option value="ETH" ${currency === 'ETH' ? 'selected' : ''}>Ξ ETH</option>
            </select>

            <div  style="grid-column:1/-1; display:flex; gap:8px; justify-content:flex-end;">
                <button type="submit" style="background:#28a745; color:white; border:none; padding:10px 16px; border-radius:6px; cursor:pointer;">
                    Сохранить
                </button>

                <button type="button"
                        onclick="htmx.trigger('#spends-table','reload')"
                        style="background:#6c757d; color:white; border:none; padding:10px 16px; border-radius:6px; cursor:pointer;">
                    Отмена
                </button>
            </div>
        </form>
    </td>
`;

    htmx.process(row);
}
