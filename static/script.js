document.addEventListener('DOMContentLoaded', () => {
    const select = document.getElementById('month_year');
    const workDataBody = document.getElementById('work-data-body');
    const baseYear = 2023;
    const baseMonth = 8;
    const startDate = new Date(baseYear, baseMonth - 1); // 2023年8月度開始

    function getMonthYearString(year, month) {
        return `${year}年 ${String(month).padStart(2, '0')}月度`;
    }

    function populateDropdown(selectedYear, selectedMonth) {
        const selectedDate = new Date(selectedYear, selectedMonth - 1);
        const startDropdownDate = new Date(selectedDate);
        startDropdownDate.setMonth(startDropdownDate.getMonth() - 11); // 選択月度の1年前まで
        const endDate = new Date(selectedDate);
        endDate.setMonth(endDate.getMonth() + 6); // 選択月度の半年後まで
        endDate.setDate(0); // 月末日

        // 開始日時が基準日時より前の場合は基準日時を開始日時に設定
        if (startDropdownDate < startDate) {
            startDropdownDate.setTime(startDate.getTime());
        }

        let currentDate = new Date(startDropdownDate);

        select.innerHTML = ''; // ドロップダウンをクリア

        while (currentDate <= endDate) {
            const year = currentDate.getFullYear();
            const month = currentDate.getMonth() + 1;
            const option = document.createElement('option');
            option.value = `${year}-${String(month).padStart(2, '0')}`;
            option.textContent = getMonthYearString(year, month);

            // 現在の年月をデフォルト選択にする
            if (year === selectedYear && month === selectedMonth) {
                option.selected = true;
            }

            select.appendChild(option);
            currentDate.setMonth(currentDate.getMonth() + 1);
        }
    }

    // 初期表示設定
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1; // JavaScriptの月は0から始まるため+1
    populateDropdown(currentYear, currentMonth);

    // ドロップダウンの選択が変更されたときにデータを更新する処理
    select.addEventListener('change', () => {
        const selectedMonthYear = select.value.split('-');
        const selectedYear = parseInt(selectedMonthYear[0]);
        const selectedMonth = parseInt(selectedMonthYear[1]);

        // ドロップダウンメニューを更新
        populateDropdown(selectedYear, selectedMonth);

        // データを更新するためのAPIリクエスト
        fetch(`/api/work_outputs?month_year=${select.value}`)
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
