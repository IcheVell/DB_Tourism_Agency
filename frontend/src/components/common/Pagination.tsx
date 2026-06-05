interface PaginationProps {
    page?: number | null;
    limit?: number | null;
    total?: number | null;
    onPageChange: (page: number) => void;
}

function normalizePositiveInteger(value: unknown, fallback: number): number {
    const parsed = Number(value);

    if (!Number.isFinite(parsed) || parsed < 1) {
        return fallback;
    }

    return Math.floor(parsed);
}

function normalizeNonNegativeInteger(value: unknown): number {
    const parsed = Number(value);

    if (!Number.isFinite(parsed) || parsed < 0) {
        return 0;
    }

    return Math.floor(parsed);
}

export function Pagination({ page, limit, total, onPageChange }: PaginationProps) {
    const currentPage = normalizePositiveInteger(page, 1);
    const currentLimit = normalizePositiveInteger(limit, 20);
    const currentTotal = normalizeNonNegativeInteger(total);

    const totalPages = Math.max(1, Math.ceil(currentTotal / currentLimit));

    return (
        <div className="pagination">
            <button
                type="button"
                disabled={currentPage <= 1}
                onClick={() => onPageChange(Math.max(1, currentPage - 1))}
            >
                Назад
            </button>

            <span>
        Страница {currentPage} из {totalPages}. Всего: {currentTotal}
      </span>

            <button
                type="button"
                disabled={currentPage >= totalPages}
                onClick={() => onPageChange(Math.min(totalPages, currentPage + 1))}
            >
                Вперёд
            </button>
        </div>
    );
}