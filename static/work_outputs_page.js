document.addEventListener('DOMContentLoaded', () => {
    const select = document.getElementById('month_year');
    const workDataBody = document.getElementById('work-data-body');
    const baseYear = 2023;
    const baseMonth = 8;
    const startDate = new Date(baseYear, baseMonth - 1); // 2023年8月度開始
    const initialMonthly = document.querySelector('meta[name="monthly"]').getAttribute('content');
    const initialMonthYear = initialMonthly.split('-');
    const initialYear = parseInt(initialMonthYear[0]);
    const initialMonth = parseInt(initialMonthYear[1]);

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
    populateDropdown(initialYear, initialMonth);

    // ドロップダウンの選択が変更されたときにデータを更新する処理
    select.addEventListener('change', () => {
        const selectedMonthYear = select.value.split('-');
        const selectedYear = parseInt(selectedMonthYear[0]);
        const selectedMonth = parseInt(selectedMonthYear[1]);

        // ドロップダウンメニューを更新
        populateDropdown(selectedYear, selectedMonth);

        // データを更新するためのページリダイレクト
        window.location.href = `/aggregate?monthly=${select.value}`;
    });
});
