import { ReactNode } from 'react';

interface DataTableProps {
    rows?: Record<string, unknown>[] | null;
    data?: Record<string, unknown>[] | null;
    columns?: string[] | null;
}

function formatCellValue(value: unknown): ReactNode {
    if (value === null || value === undefined) {
        return '—';
    }

    if (typeof value === 'boolean') {
        return value ? 'Да' : 'Нет';
    }

    if (
        typeof value === 'object'
        && 'type' in value
        && 'props' in value
    ) {
        return value as ReactNode;
    }

    if (typeof value === 'object') {
        return JSON.stringify(value);
    }

    return String(value);
}

export function DataTable({ rows, data, columns }: DataTableProps) {
    const items = Array.isArray(rows)
        ? rows
        : Array.isArray(data)
            ? data
            : [];

    const tableColumns = Array.isArray(columns) && columns.length > 0
        ? columns
        : items.length > 0
            ? Object.keys(items[0])
            : [];

    if (items.length === 0) {
        return <div className="empty-state">Нет данных</div>;
    }

    return (
        <div className="table-wrapper">
            <table className="data-table">
                <thead>
                <tr>
                    {tableColumns.map((column) => (
                        <th key={column}>{column}</th>
                    ))}
                </tr>
                </thead>

                <tbody>
                {items.map((item, rowIndex) => (
                    <tr key={rowIndex}>
                        {tableColumns.map((column) => (
                            <td key={column}>{formatCellValue(item[column])}</td>
                        ))}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
}