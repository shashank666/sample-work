case parser.LIMITK:
    param, err := evaluateScalar(aggr.Param, step)
    if err != nil {
        return nil, err
    }
    k := int(param)
    if k < 0 {
        return nil, errors.New("limitk parameter must be non-negative")
    }

    for _, group := range groupedSeries {
        seriesWithHashes := make([]seriesWithHash, 0, len(group))
        for _, s := range group {
            hash := s.Labels().Hash()
            seriesWithHashes = append(seriesWithHashes, seriesWithHash{s, hash})
        }

      
        sort.Slice(seriesWithHashes, func(i, j int) bool {
            return seriesWithHashes[i].hash < seriesWithHashes[j].hash
        })

       
        limit := k
        if limit > len(seriesWithHashes) {
            limit = len(seriesWithHashes)
        }

      
        for i := 0; i < limit; i++ {
            result = append(result, seriesWithHashes[i].series)
        }
    }
