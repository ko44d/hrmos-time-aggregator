document.addEventListener('DOMContentLoaded', () => {
    const select = document.getElementById('month_year');
    const workDataBody = document.getElementById('work-data-body');
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1; // JavaScriptの月は0から始まるため+1
    const startYear = currentYear - 1;
    const endYear = currentYear + 1;

    for (let year = startYear; year <= endYear; year++) {
        for (let month = 1; month <= 12; month++) {
            const monthYear = `${year}年 ${String(month).padStart(2, '0')}月度`;
            const option = document.createElement('option');
            option.value = `${year}-${String(month).padStart(2, '0')}`;
            option.textContent = monthYear;

            // 現在の年月をデフォルト選択にする
            if (year === currentYear && month === currentMonth) {
                option.selected = true;
            }

            select.appendChild(option);
        }
    }

    // ドロップダウンの選択が変更されたときにデータを更新する処理
    select.addEventListener('change', () => {
        const selectedMonthYear = select.value;

        fetch(`/api/work_outputs?month_year=${selectedMonthYear}`)
            .then(response => response.json())
            .then(data => {
                // テーブルの内容をクリア
                workDataBody.innerHTML = '';

                // 新しいデータでテーブルを更新
                data.forEach(item => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${item.UserId}</td>
                        <td>${item.Number}</td>
                        <td>${item.FullName}</td>
                        <td>${item.Month}</td>
                        <td>${item.Day}</td>
                        <td>${item.Wday}</td>
                        <td>${item.StartAt}</td>
                        <td>${item.EndAt}</td>
                        <td>${item.TotalOverWorkTime}</td>
                    `;
                    workDataBody.appendChild(row);
                });
            })
            .catch(error => {
                console.error('Error fetching work outputs data:', error);
            });
    });

    // 初期データの取得
    select.dispatchEvent(new Event('change'));
});
