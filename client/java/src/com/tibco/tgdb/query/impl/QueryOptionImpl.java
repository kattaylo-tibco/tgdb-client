package com.tibco.tgdb.query.impl;
/**
 * Copyright 2016 TIBCO Software Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); You may not use this file except
 * in compliance with the License.
 * A copy of the License is included in the distribution package with this file.
 * You also may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File name : TGQueryOption.java
 * SVN Id: $Id$
 */

import com.tibco.tgdb.query.TGQueryOption;
import com.tibco.tgdb.utils.SortedProperties;


public class QueryOptionImpl extends SortedProperties<String,String> implements TGQueryOption {

    private boolean mutable;
    private static String QUERY_OPTION_FETCHSIZE = "fetchsize";
    private static String QUERY_OPTION_TRAVERSALDEPTH = "traversaldepth";
    private static String QUERY_OPTION_EDGELIMIT = "edgelimit";


    public QueryOptionImpl(boolean mutable) {
        this.mutable = mutable;
        this.put(QUERY_OPTION_FETCHSIZE, "-1");
        this.put(QUERY_OPTION_TRAVERSALDEPTH, "-1");
        this.put(QUERY_OPTION_EDGELIMIT, "-1");

    }

    @Override
    public void setPrefetchSize(int size) {
        if (!mutable) throw new RuntimeException("Can't modify a immutable Option");
        if (size == 0) size = -1;
        this.put(QUERY_OPTION_FETCHSIZE, String.valueOf(size));
    }

    @Override
    public int getPrefetchSize() {
        String s = this.get(QUERY_OPTION_FETCHSIZE);
        return Integer.valueOf(s);
    }

    @Override
    public void setTraversalDepth(int depth) {
        if (!mutable) throw new RuntimeException("Can't modify a immutable Option");
        if (depth == 0) depth = -1;
        this.put(QUERY_OPTION_TRAVERSALDEPTH, String.valueOf(depth));
    }

    @Override
    public int getTraversalDepth() {
        String s = this.get(QUERY_OPTION_TRAVERSALDEPTH);
        return Integer.valueOf(s);
    }

    @Override
    public void setEdgeLimit(int limit) {
        if (!mutable) throw new RuntimeException("Can't modify a immutable Option");
        if (limit == 0) limit = -1;
        this.put(QUERY_OPTION_EDGELIMIT, String.valueOf(limit));
    }

    @Override
    public int getEdgeLimit() {
        String s = this.get(QUERY_OPTION_EDGELIMIT);
        return Integer.valueOf(s);
    }
}
