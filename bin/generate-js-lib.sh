#!/bin/bash
if [ $# -ne 3 ]
then
    echo "Error in $0 - Invalid Argument Count "
    echo "Syntax: $0 location_of_service api_key library_root"
    exit
fi

echo "" > classpath.txt
for file in `ls lib`;
        do echo -n 'lib/' >> classpath.txt;
        echo -n $file >> classpath.txt;
        echo -n ':' >> classpath.txt;
done
for file in `ls build`;
	do echo -n 'build/' >> classpath.txt;
	echo -n $file >> classpath.txt;
	echo -n ':' >> classpath.txt;
done

export CLASSPATH=$(cat classpath.txt):conf/js/templates
export JAVA_OPTS="${JAVA_OPTS} -Dproperty=Xmx2g"
java $WORDNIK_OPTS $JAVA_CONFIG_OPTIONS $JAVA_OPTS -cp $CLASSPATH com.wordnik.swagger.codegen.config.js.JSLibCodeGen "$@"