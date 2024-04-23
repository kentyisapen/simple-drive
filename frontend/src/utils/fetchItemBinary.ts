type Result =
    | {
        data: ArrayBuffer;
        error: null;
    }
    | {
        data: null;
        error: Error;
    };

export const fetchItemBinary = async (itemId: string): Promise<Result> => {
    try {
        const response = await fetch(process.env.NEXT_PUBLIC_BASE_MINIO_OBJECT_URL + itemId)
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.arrayBuffer();
        return { data, error: null };
    } catch (e) {
        return { data: null, error: e instanceof Error ? e : new Error(String(e)) };
    }
}